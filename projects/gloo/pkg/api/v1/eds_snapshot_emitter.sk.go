// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"sync"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	"go.uber.org/zap"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/errors"
	skstats "github.com/solo-io/solo-kit/pkg/stats"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/errutils"
)

var (
	// Deprecated. See mEdsResourcesIn
	mEdsSnapshotIn = stats.Int64("eds.gloo.solo.io/emitter/snap_in", "Deprecated. Use eds.gloo.solo.io/emitter/resources_in. The number of snapshots in", "1")

	// metrics for emitter
	mEdsResourcesIn    = stats.Int64("eds.gloo.solo.io/emitter/resources_in", "The number of resource lists received on open watch channels", "1")
	mEdsSnapshotOut    = stats.Int64("eds.gloo.solo.io/emitter/snap_out", "The number of snapshots out", "1")
	mEdsSnapshotMissed = stats.Int64("eds.gloo.solo.io/emitter/snap_missed", "The number of snapshots missed", "1")

	// views for emitter
	// deprecated: see edsResourcesInView
	edssnapshotInView = &view.View{
		Name:        "eds.gloo.solo.io/emitter/snap_in",
		Measure:     mEdsSnapshotIn,
		Description: "Deprecated. Use eds.gloo.solo.io/emitter/resources_in. The number of snapshots updates coming in.",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}

	edsResourcesInView = &view.View{
		Name:        "eds.gloo.solo.io/emitter/resources_in",
		Measure:     mEdsResourcesIn,
		Description: "The number of resource lists received on open watch channels",
		Aggregation: view.Count(),
		TagKeys: []tag.Key{
			skstats.NamespaceKey,
			skstats.ResourceKey,
		},
	}
	edssnapshotOutView = &view.View{
		Name:        "eds.gloo.solo.io/emitter/snap_out",
		Measure:     mEdsSnapshotOut,
		Description: "The number of snapshots updates going out",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
	edssnapshotMissedView = &view.View{
		Name:        "eds.gloo.solo.io/emitter/snap_missed",
		Measure:     mEdsSnapshotMissed,
		Description: "The number of snapshots updates going missed. this can happen in heavy load. missed snapshot will be re-tried after a second.",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{},
	}
)

func init() {
	view.Register(
		edssnapshotInView,
		edssnapshotOutView,
		edssnapshotMissedView,
		edsResourcesInView,
	)
}

type EdsSnapshotEmitter interface {
	Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *EdsSnapshot, <-chan error, error)
}

type EdsEmitter interface {
	EdsSnapshotEmitter
	Register() error
	Upstream() UpstreamClient
}

func NewEdsEmitter(upstreamClient UpstreamClient) EdsEmitter {
	return NewEdsEmitterWithEmit(upstreamClient, make(chan struct{}))
}

func NewEdsEmitterWithEmit(upstreamClient UpstreamClient, emit <-chan struct{}) EdsEmitter {
	return &edsEmitter{
		upstream:  upstreamClient,
		forceEmit: emit,
	}
}

type edsEmitter struct {
	forceEmit <-chan struct{}
	upstream  UpstreamClient
}

func (c *edsEmitter) Register() error {
	if err := c.upstream.Register(); err != nil {
		return err
	}
	return nil
}

func (c *edsEmitter) Upstream() UpstreamClient {
	return c.upstream
}

func (c *edsEmitter) Snapshots(watchNamespaces []string, opts clients.WatchOpts) (<-chan *EdsSnapshot, <-chan error, error) {

	if len(watchNamespaces) == 0 {
		watchNamespaces = []string{""}
	}

	for _, ns := range watchNamespaces {
		if ns == "" && len(watchNamespaces) > 1 {
			return nil, nil, errors.Errorf("the \"\" namespace is used to watch all namespaces. Snapshots can either be tracked for " +
				"specific namespaces or \"\" AllNamespaces, but not both.")
		}
	}

	errs := make(chan error)
	var done sync.WaitGroup
	ctx := opts.Ctx
	/* Create channel for Upstream */
	type upstreamListWithNamespace struct {
		list      UpstreamList
		namespace string
	}
	upstreamChan := make(chan upstreamListWithNamespace)

	var initialUpstreamList UpstreamList

	currentSnapshot := EdsSnapshot{}

	for _, namespace := range watchNamespaces {
		/* Setup namespaced watch for Upstream */
		{
			upstreams, err := c.upstream.List(namespace, clients.ListOpts{Ctx: opts.Ctx, Selector: opts.Selector})
			if err != nil {
				return nil, nil, errors.Wrapf(err, "initial Upstream list")
			}
			initialUpstreamList = append(initialUpstreamList, upstreams...)
		}
		upstreamNamespacesChan, upstreamErrs, err := c.upstream.Watch(namespace, opts)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "starting Upstream watch")
		}

		done.Add(1)
		go func(namespace string) {
			defer done.Done()
			errutils.AggregateErrs(ctx, errs, upstreamErrs, namespace+"-upstreams")
		}(namespace)

		/* Watch for changes and update snapshot */
		go func(namespace string) {
			for {
				select {
				case <-ctx.Done():
					return
				case upstreamList := <-upstreamNamespacesChan:
					select {
					case <-ctx.Done():
						return
					case upstreamChan <- upstreamListWithNamespace{list: upstreamList, namespace: namespace}:
					}
				}
			}
		}(namespace)
	}
	/* Initialize snapshot for Upstreams */
	currentSnapshot.Upstreams = initialUpstreamList.Sort()

	snapshots := make(chan *EdsSnapshot)
	go func() {
		// sent initial snapshot to kick off the watch
		initialSnapshot := currentSnapshot.Clone()
		snapshots <- &initialSnapshot

		timer := time.NewTicker(time.Second * 1)
		previousHash, err := currentSnapshot.Hash(nil)
		if err != nil {
			contextutils.LoggerFrom(ctx).DPanicw("error while hashing, this should never happen", zap.Error(err))
		}
		sync := func() {
			currentHash, err := currentSnapshot.Hash(nil)
			// this should never happen, so panic if it does
			if err != nil {
				contextutils.LoggerFrom(ctx).DPanicw("error while hashing, this should never happen", zap.Error(err))
			}
			if previousHash == currentHash {
				return
			}

			sentSnapshot := currentSnapshot.Clone()
			select {
			case snapshots <- &sentSnapshot:
				stats.Record(ctx, mEdsSnapshotOut.M(1))
				previousHash = currentHash
			default:
				stats.Record(ctx, mEdsSnapshotMissed.M(1))
			}
		}
		upstreamsByNamespace := make(map[string]UpstreamList)

		for {
			record := func() { stats.Record(ctx, mEdsSnapshotIn.M(1)) }

			select {
			case <-timer.C:
				sync()
			case <-ctx.Done():
				close(snapshots)
				done.Wait()
				close(errs)
				return
			case <-c.forceEmit:
				sentSnapshot := currentSnapshot.Clone()
				snapshots <- &sentSnapshot
			case upstreamNamespacedList := <-upstreamChan:
				record()

				namespace := upstreamNamespacedList.namespace

				skstats.IncrementResourceCount(
					ctx,
					namespace,
					"upstream",
					mEdsResourcesIn,
				)

				// merge lists by namespace
				upstreamsByNamespace[namespace] = upstreamNamespacedList.list
				var upstreamList UpstreamList
				for _, upstreams := range upstreamsByNamespace {
					upstreamList = append(upstreamList, upstreams...)
				}
				currentSnapshot.Upstreams = upstreamList.Sort()
			}
		}
	}()
	return snapshots, errs, nil
}
