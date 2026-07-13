## Profile

```
   -17.69s 21.78% 21.78%    -17.68s 21.77%  crypto/internal/fips140/sha256.blockAVX2
     6.94s  8.54% 13.24%      6.94s  8.54%  internal/runtime/syscall/linux.Syscall6
    -1.45s  1.79% 15.02%    -20.49s 25.23%  crypto/internal/fips140/sha256.(*Digest).Write
    -1.25s  1.54% 16.56%    -21.59s 26.58%  crypto/internal/fips140/sha256.(*Digest).Sum
    -1.12s  1.38% 17.94%    -25.71s 31.65%  crypto/internal/fips140/pbkdf2.Key[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }]
    -0.82s  1.01% 18.95%    -23.53s 28.97%  crypto/internal/fips140/hmac.(*HMAC).Sum
    -0.70s  0.86% 19.81%     -0.70s  0.86%  runtime.memmove
    -0.58s  0.71% 20.52%     -1.04s  1.28%  crypto/internal/fips140/sha256.(*Digest).UnmarshalBinary
     0.53s  0.65% 19.87%      1.24s  1.53%  time.Time.appendFormat
    -0.52s  0.64% 20.51%    -20.10s 24.75%  crypto/internal/fips140/sha256.(*Digest).checkSum
    -0.44s  0.54% 21.05%    -18.14s 22.33%  crypto/internal/fips140/sha256.block
     0.44s  0.54% 20.51%      0.45s  0.55%  runtime.(*mspan).writeHeapBitsSmall
     0.43s  0.53% 19.98%      0.69s  0.85%  go.uber.org/zap/zapcore.safeAppendStringLike[go.shape.string]
     0.43s  0.53% 19.45%      1.15s  1.42%  runtime.pcvalue
     0.42s  0.52% 18.94%      0.71s  0.87%  runtime.tryDeferToSpanScan
     0.40s  0.49% 18.44%      0.49s   0.6%  runtime.step
     0.36s  0.44% 18.00%      0.36s  0.44%  runtime.nextFreeFast (inline)
     0.30s  0.37% 17.63%      0.40s  0.49%  net/textproto.CanonicalMIMEHeaderKey
     0.30s  0.37% 17.26%      0.31s  0.38%  time.nextStdChunk
     0.28s  0.34% 16.92%      0.32s  0.39%  runtime.findfunc
     0.27s  0.33% 16.58%      0.27s  0.33%  indexbytebody
     0.24s   0.3% 16.29%      0.24s   0.3%  time.runtimeNow
     0.22s  0.27% 16.02%      0.36s  0.44%  go.uber.org/zap/buffer.(*Buffer).AppendString (partial-inline)
     0.21s  0.26% 15.76%      0.21s  0.26%  runtime.asyncPreempt
     0.20s  0.25% 15.51%      0.36s  0.44%  runtime.scanObjectsSmall
     0.19s  0.23% 15.28%      1.23s  1.51%  go.uber.org/zap/zapcore.Field.AddTo
     0.19s  0.23% 15.05%      0.19s  0.23%  runtime.spanClass.sizeclass (inline)
     0.18s  0.22% 14.82%      1.83s  2.25%  net/http.(*chunkWriter).writeHeader
     0.18s  0.22% 14.60%      0.59s  0.73%  runtime.(*unwinder).resolveInternal
     0.18s  0.22% 14.38%      0.18s  0.22%  runtime.madvise
    -0.17s  0.21% 14.59%     -0.15s  0.18%  crypto/internal/fips140.RecordApproved
     0.17s  0.21% 14.38%      1.75s  2.15%  runtime.mallocgcSmallScanNoHeader
     0.17s  0.21% 14.17%      0.45s  0.55%  runtime.scanObject
     0.17s  0.21% 13.96%      0.17s  0.21%  time.appendInt
    -0.16s   0.2% 14.16%     -0.34s  0.42%  crypto/internal/fips140/sha256.consumeUint32 (inline)
     0.15s  0.18% 13.97%      0.23s  0.28%  runtime.findObject
     0.15s  0.18% 13.79%      0.44s  0.54%  sync.(*Pool).Get
     0.14s  0.17% 13.62%      0.14s  0.17%  runtime.memclrNoHeapPointers
     0.13s  0.16% 13.46%      0.13s  0.16%  aeshashbody
     0.13s  0.16% 13.30%      0.16s   0.2%  context.value
     0.13s  0.16% 13.14%      2.46s  3.03%  go.uber.org/zap.(*Logger).check
     0.13s  0.16% 12.98%      3.82s  4.70%  go.uber.org/zap/zapcore.consoleEncoder.EncodeEntry
     0.13s  0.16% 12.82%      0.61s  0.75%  internal/poll.runtime_pollSetDeadline
     0.13s  0.16% 12.66%      0.13s  0.16%  runtime.procyieldAsm
     0.13s  0.16% 12.50%      0.13s  0.16%  runtime.releasem (inline)
     0.13s  0.16% 12.34%      0.18s  0.22%  runtime.scanblock
     0.13s  0.16% 12.18%      0.95s  1.17%  runtime.tracebackPCs
     0.13s  0.16% 12.02%      0.12s  0.15%  runtime.typePointers.next
     0.13s  0.16% 11.86%      0.13s  0.16%  runtime.usleep
     0.13s  0.16% 11.70%      0.78s  0.96%  runtime.wbBufFlush1
    -0.12s  0.15% 11.84%     -0.77s  0.95%  crypto/internal/fips140/hmac.(*HMAC).Write (inline)
    -0.12s  0.15% 11.99%     -0.13s  0.16%  internal/byteorder.BEUint32 (inline)
     0.12s  0.15% 11.84%      0.33s  0.41%  internal/sync.(*Mutex).Lock (inline)
     0.12s  0.15% 11.70%      0.77s  0.95%  net/http.(*conn).serve
     0.12s  0.15% 11.55%      0.26s  0.32%  runtime.lock2
     0.12s  0.15% 11.40%      2.90s  3.57%  runtime.mallocgc
     0.12s  0.15% 11.25%      0.22s  0.27%  runtime.mapaccess2_faststr
     0.12s  0.15% 11.11%      0.36s  0.44%  sync.(*Pool).Put
     0.11s  0.14% 10.97%      0.22s  0.27%  github.com/go-chi/chi/v5.(*node).findRoute
     0.11s  0.14% 10.83%      0.11s  0.14%  go.uber.org/zap/buffer.(*Buffer).AppendByte (inline)
     0.11s  0.14% 10.70%    -12.98s 15.98%  main.main.WithLogger.func6.1
     0.11s  0.14% 10.56%      2.91s  3.58%  net/http.(*conn).readRequest
     0.11s  0.14% 10.43%      0.39s  0.48%  runtime.(*stkframe).getStackMap
     0.10s  0.12% 10.31%      0.12s  0.15%  internal/bytealg.LastIndexByteString (inline)
    -0.10s  0.12% 10.43%     -0.11s  0.14%  internal/byteorder.BEPutUint32 (inline)
     0.10s  0.12% 10.31%      0.10s  0.12%  internal/runtime/atomic.(*Uint32).CompareAndSwap (inline)
     0.10s  0.12% 10.18%      0.10s  0.12%  net/textproto.validHeaderFieldByte (inline)
     0.10s  0.12% 10.06%      0.10s  0.12%  runtime.acquireSudog
     0.10s  0.12%  9.94%      0.10s  0.12%  runtime.nanotime (inline)
     0.10s  0.12%  9.81%      0.10s  0.12%  runtime.scanObjectSmall
     0.10s  0.12%  9.69%      0.34s  0.42%  time.Now
     0.09s  0.11%  9.58%      6.33s  7.79%  go.uber.org/zap/zapcore.(*CheckedEntry).Write
     0.09s  0.11%  9.47%      0.09s  0.11%  runtime.(*moduledata).textAddr
     0.09s  0.11%  9.36%      0.09s  0.11%  runtime.acquirem (inline)
    -0.09s  0.11%  9.47%     -0.09s  0.11%  runtime.futex
     0.09s  0.11%  9.36%      0.26s  0.32%  runtime.mapaccess1_faststr
     0.09s  0.11%  9.25%      0.09s  0.11%  runtime.readvarint (inline)
     0.09s  0.11%  9.14%      0.16s   0.2%  runtime.unlock2
     0.09s  0.11%  9.02%      0.13s  0.16%  sync.(*Pool).pin
     0.08s 0.098%  8.93%      0.38s  0.47%  encoding/json.structEncoder.encode
     0.08s 0.098%  8.83%      1.84s  2.27%  net/http.readRequest
     0.08s 0.098%  8.73%      0.63s  0.78%  runtime.(*Frames).Next
     0.08s 0.098%  8.63%      0.14s  0.17%  runtime.ifaceeq
     0.08s 0.098%  8.53%      0.19s  0.23%  runtime.mallocgcSmallNoscan
     0.08s 0.098%  8.43%      1.64s  2.02%  runtime.newobject
     0.07s 0.086%  8.35%      0.16s   0.2%  bufio.(*Writer).WriteString
     0.07s 0.086%  8.26%      0.07s 0.086%  github.com/go-chi/chi/v5.nodes.findEdge (inline)
     0.07s 0.086%  8.18%    -12.96s 15.96%  net/http.HandlerFunc.ServeHTTP
     0.07s 0.086%  8.09%      0.45s  0.55%  net/textproto.MIMEHeader.Get
     0.07s 0.086%  8.00%      0.41s   0.5%  runtime.(*mcache).refill
     0.07s 0.086%  7.92%      0.13s  0.16%  runtime.concatstrings
     0.07s 0.086%  7.83%      0.18s  0.22%  runtime.makeslice
     0.07s 0.086%  7.74%      0.10s  0.12%  runtime.mallocgcTiny
     0.07s 0.086%  7.66%      0.07s 0.086%  runtime.nilinterhash
     0.07s 0.086%  7.57%      0.34s  0.42%  runtime.park_m
     0.07s 0.086%  7.49%      0.10s  0.12%  runtime.rand
     0.07s 0.086%  7.40%      0.07s 0.086%  runtime.runqput
     0.07s 0.086%  7.31%      0.44s  0.54%  runtime.slicebytetostring
     0.06s 0.074%  7.24%      0.21s  0.26%  context.WithValue
     0.06s 0.074%  7.17%      0.65s   0.8%  encoding/json.(*encodeState).reflectValue
     0.06s 0.074%  7.09%      0.14s  0.17%  fmt.(*pp).printArg
     0.06s 0.074%  7.02%      2.18s  2.68%  go.uber.org/zap/zapcore.(*lockedWriteSyncer).Write
     0.06s 0.074%  6.94%      1.29s  1.59%  go.uber.org/zap/zapcore.addFields (inline)
     0.06s 0.074%  6.87%      0.06s 0.074%  internal/poll.(*fdMutex).decref (inline)
     0.06s 0.074%  6.80%         1s  1.23%  internal/poll.setDeadlineImpl
     0.06s 0.074%  6.72%      0.16s   0.2%  internal/runtime/maps.(*Iter).Next
     0.06s 0.074%  6.65%      0.06s 0.074%  internal/runtime/maps.(*ctrlGroup).setEmpty (inline)
     0.06s 0.074%  6.57%      0.06s 0.074%  internal/strconv.formatBase10
     0.06s 0.074%  6.50%      0.06s 0.074%  memeqbody
     0.06s 0.074%  6.43%      0.12s  0.15%  net/http.(*Request).WithContext (inline)
     0.06s 0.074%  6.35%      0.06s 0.074%  net/http.(*response).Header
     0.06s 0.074%  6.28%      0.65s   0.8%  net/http.Header.writeSubset
     0.06s 0.074%  6.21%      0.79s  0.97%  net/textproto.readMIMEHeader
     0.06s 0.074%  6.13%      0.07s 0.086%  net/url.unescape
     0.06s 0.074%  6.06%      0.06s 0.074%  reflect.Value.Field
     0.06s 0.074%  5.98%      0.07s 0.086%  runtime.gopark
     0.06s 0.074%  5.91%      0.08s 0.098%  runtime.growslice
     0.06s 0.074%  5.84%      0.15s  0.18%  runtime.reentersyscall
     0.06s 0.074%  5.76%      0.06s 0.074%  runtime.strhash
     0.06s 0.074%  5.69%      0.16s   0.2%  strings.(*Replacer).Replace
     0.06s 0.074%  5.61%      0.06s 0.074%  strings.(*byteReplacer).Replace
     0.06s 0.074%  5.54%      0.06s 0.074%  time.Time.Add
     0.06s 0.074%  5.47%      0.08s 0.098%  time.Time.locabs
     0.06s 0.074%  5.39%      0.21s  0.26%  time.Until
     0.06s 0.074%  5.32%      0.06s 0.074%  vendor/golang.org/x/net/http/httpguts.ValidHostHeader (inline)
     0.05s 0.062%  5.26%         7s  8.62%  bufio.(*Writer).Flush
    -0.05s 0.062%  5.32%     -0.18s  0.22%  crypto/internal/fips140deps/byteorder.BEUint32 (inline)
     0.05s 0.062%  5.26%      0.32s  0.39%  fmt.Fprint
     0.05s 0.062%  5.20%      0.54s  0.66%  gcWriteBarrier
     0.05s 0.062%  5.13%    -12.61s 15.53%  github.com/go-chi/chi/v5.(*Mux).ServeHTTP
     0.05s 0.062%  5.07%    -19.09s 23.50%  github.com/paulwwyvern/urlshortener/pkg/httphelpers/httperr.Adapt.func1
     0.05s 0.062%  5.01%      6.15s  7.57%  go.uber.org/zap/zapcore.(*ioCore).Write
     0.05s 0.062%  4.95%      0.09s  0.11%  go.uber.org/zap/zapcore.(*jsonEncoder).clone
     0.05s 0.062%  4.89%      0.29s  0.36%  go.uber.org/zap/zapcore.(*sliceArrayEncoder).AppendString
     0.05s 0.062%  4.83%      6.20s  7.63%  internal/poll.(*FD).Write
     0.05s 0.062%  4.76%      0.05s 0.062%  internal/runtime/atomic.(*Uint32).Add (inline)
     0.05s 0.062%  4.70%      0.11s  0.14%  internal/runtime/maps.(*Iter).Init
     0.05s 0.062%  4.64%      0.28s  0.34%  internal/runtime/maps.(*Map).Delete
     0.05s 0.062%  4.58%      0.05s 0.062%  internal/runtime/maps.ctrlGroup.matchH2 (inline)
     0.05s 0.062%  4.52%      0.17s  0.21%  internal/strconv.AppendUint
     0.05s 0.062%  4.46%      2.10s  2.59%  net.(*conn).Read
     0.05s 0.062%  4.40%      1.54s  1.90%  net/http.Redirect
     0.05s 0.062%  4.33%      0.05s 0.062%  runtime.(*mLockProfile).store (inline)
     0.05s 0.062%  4.27%      0.08s 0.098%  runtime.(*timer).stop
     0.05s 0.062%  4.21%      0.09s  0.11%  runtime.execute
     0.05s 0.062%  4.15%      0.06s 0.074%  runtime.heapSetTypeSmallHeader (inline)
     0.05s 0.062%  4.09%      0.81s     1%  runtime.markroot
     0.05s 0.062%  4.03%      0.05s 0.062%  runtime.osyield
     0.05s 0.062%  3.96%      0.54s  0.66%  runtime.scanSpan
     0.05s 0.062%  3.90%      0.14s  0.17%  slices.insertionSortCmpFunc[go.shape.struct { net/http.key string; net/http.values []string }] (inline)
     0.05s 0.062%  3.84%      0.05s 0.062%  strings.ToLower
     0.05s 0.062%  3.78%      0.16s   0.2%  sync.(*poolChain).pushHead
     0.05s 0.062%  3.72%      0.05s 0.062%  time.now
     0.04s 0.049%  3.67%      0.04s 0.049%  bytes.(*Buffer).Write
     0.04s 0.049%  3.62%      0.04s 0.049%  bytes.makeASCIISet (inline)
     0.04s 0.049%  3.57%      0.05s 0.062%  context.(*cancelCtx).Done
    -0.04s 0.049%  3.62%     -0.38s  0.47%  crypto/internal/fips140/hmac.(*HMAC).Reset
    -0.04s 0.049%  3.67%     -0.04s 0.049%  crypto/internal/fips140deps/byteorder.BEPutUint64 (inline)
     0.04s 0.049%  3.62%      0.04s 0.049%  encoding/json.appendString[go.shape.string]
     0.04s 0.049%  3.57%     -0.09s  0.11%  github.com/jackc/pgx/v5/pgconn.(*PgConn).receiveMessage
     0.04s 0.049%  3.52%    -24.81s 30.55%  github.com/paulwwyvern/urlshortener/internal/handler/chihttp.(*Handler).GetURL.Adapt.func1 (inline)
     0.04s 0.049%  3.47%    -24.89s 30.65%  github.com/paulwwyvern/urlshortener/internal/handler/chihttp.(*Handler).getURL
     0.04s 0.049%  3.42%      5.23s  6.44%  github.com/paulwwyvern/urlshortener/internal/repository/audit.(*LogLogger).Update
     0.04s 0.049%  3.37%      0.22s  0.27%  github.com/paulwwyvern/urlshortener/pkg/httphelpers/httpurl.SetURL
     0.04s 0.049%  3.32%      0.15s  0.18%  go.uber.org/zap/zapcore.CapitalLevelEncoder
     0.04s 0.049%  3.28%      1.16s  1.43%  go.uber.org/zap/zapcore.encodeTimeLayout
     0.04s 0.049%  3.23%      0.20s  0.25%  go.uber.org/zap/zapcore.getCheckedEntry
    -0.04s 0.049%  3.28%      1.57s  1.93%  internal/poll.(*FD).Read
     0.04s 0.049%  3.23%      0.04s 0.049%  internal/poll.(*fdMutex).incref (inline)
     0.04s 0.049%  3.18%      0.04s 0.049%  internal/runtime/atomic.(*Uint8).Load (inline)
     0.04s 0.049%  3.13%      0.04s 0.049%  internal/runtime/maps.(*Map).Used (inline)
     0.04s 0.049%  3.08%      0.16s   0.2%  internal/runtime/maps.(*Map).getWithoutKeySmallFastStr
     0.04s 0.049%  3.03%      0.05s 0.062%  internal/runtime/maps.(*Map).putSlotSmallFastStr
     0.04s 0.049%  2.98%      1.40s  1.72%  net/http.(*connReader).Read
     0.04s 0.049%  2.93%      7.56s  9.31%  net/http.(*response).finishRequest
     0.04s 0.049%  2.88%      0.10s  0.12%  net/http.Header.get (inline)
     0.04s 0.049%  2.83%      0.34s  0.42%  net/http.Header.sortedKeyValues
     0.04s 0.049%  2.78%      0.09s  0.11%  net/http.newTextprotoReader
     0.04s 0.049%  2.73%      0.13s  0.16%  net/http.readTransfer
     0.04s 0.049%  2.68%      0.36s  0.44%  net/url.Parse
     0.04s 0.049%  2.63%      0.04s 0.049%  reflect.Value.Elem
     0.04s 0.049%  2.59%      0.04s 0.049%  runtime.(*mspan).inlineMarkBits (inline)
     0.04s 0.049%  2.54%      0.04s 0.049%  runtime.(*pallocBits).summarize
     0.04s 0.049%  2.49%      0.06s 0.074%  runtime.(*timer).unlock (inline)
     0.04s 0.049%  2.44%      0.44s  0.54%  runtime.convTstring
     0.04s 0.049%  2.39%      0.07s 0.086%  runtime.efaceeq
     0.04s 0.049%  2.34%      0.06s 0.074%  runtime.exitsyscall
     0.04s 0.049%  2.29%      0.17s  0.21%  runtime.findRunnable
     0.04s 0.049%  2.24%      0.14s  0.17%  runtime.gcmarknewobject
     0.04s 0.049%  2.19%      0.04s 0.049%  runtime.getpid
     0.04s 0.049%  2.14%      0.05s 0.062%  runtime.greyobject
     0.04s 0.049%  2.09%      0.09s  0.11%  runtime.interhash
    -0.04s 0.049%  2.14%      0.03s 0.037%  runtime.mapassign
     0.04s 0.049%  2.09%      0.62s  0.76%  runtime.mapassign_faststr
     0.04s 0.049%  2.04%      0.32s  0.39%  runtime.schedule
     0.04s 0.049%  1.99%      0.04s 0.049%  runtime.spanOf (inline)
     0.04s 0.049%  1.95%      4.57s  5.63%  runtime.systemstack
     0.04s 0.049%  1.90%      0.06s 0.074%  sync.(*poolDequeue).popHead
     0.04s 0.049%  1.85%      0.04s 0.049%  sync.(*poolDequeue).pushHead
     0.04s 0.049%  1.80%      0.08s 0.098%  sync/atomic.StorePointer
     0.04s 0.049%  1.75%      0.04s 0.049%  time.Time.Equal
     0.04s 0.049%  1.70%      0.04s 0.049%  time.Weekday.String (inline)
     0.03s 0.037%  1.66%      0.10s  0.12%  bufio.(*Writer).Write
     0.03s 0.037%  1.63%      0.03s 0.037%  crypto/internal/fips140/hmac.New[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }]
     0.03s 0.037%  1.59%      0.82s  1.01%  encoding/json.Marshal
     0.03s 0.037%  1.55%      0.05s 0.062%  encoding/json.newEncodeState
     0.03s 0.037%  1.51%      0.46s  0.57%  encoding/json.ptrEncoder.encode
     0.03s 0.037%  1.48%      0.07s 0.086%  fmt.(*fmt).fmtS
     0.03s 0.037%  1.44%      0.28s  0.34%  github.com/go-chi/chi/v5.(*node).FindRoute
    -0.03s 0.037%  1.48%     -0.13s  0.16%  github.com/jackc/pgx/v5/pgproto3.(*Frontend).Receive
     0.03s 0.037%  1.44%     49.33s 60.74%  github.com/paulwwyvern/urlshortener/internal/repository/storage/throughcache.(*Cache).GetURL
     0.03s 0.037%  1.40%      5.26s  6.48%  github.com/paulwwyvern/urlshortener/internal/service/audit.(*Publisher).Notify
     0.03s 0.037%  1.37%      0.03s 0.037%  go.uber.org/multierr.Append
     0.03s 0.037%  1.33%      0.24s   0.3%  go.uber.org/zap/internal/pool.(*Pool[go.shape.*uint8]).Put (inline)
     0.03s 0.037%  1.29%      0.66s  0.81%  go.uber.org/zap/internal/stacktrace.(*Stack).Next (inline)
     0.03s 0.037%  1.26%      0.04s 0.049%  go.uber.org/zap/zapcore.(*jsonEncoder).addElementSeparator (inline)
     0.03s 0.037%  1.22%      0.42s  0.52%  go.uber.org/zap/zapcore.(*jsonEncoder).addKey
     0.03s 0.037%  1.18%      0.03s 0.037%  internal/runtime/atomic.(*Uint64).Add (inline)
     0.03s 0.037%  1.15%      0.10s  0.12%  internal/runtime/maps.(*Map).deleteSmall
     0.03s 0.037%  1.11%      0.06s 0.074%  internal/runtime/maps.typedmemclr
     0.03s 0.037%  1.07%      0.07s 0.086%  internal/stringslite.Cut
     0.03s 0.037%  1.03%      0.11s  0.14%  internal/sync.(*HashTrieMap[go.shape.interface {},go.shape.interface {}]).Load
     0.03s 0.037%     1%      0.20s  0.25%  internal/sync.(*Mutex).lockSlow
     0.03s 0.037%  0.96%      4.50s  5.54%  net.(*netFD).Write
     0.03s 0.037%  0.92%      5.07s  6.24%  net/http.checkConnErrorWriter.Write
     0.03s 0.037%  0.89%      0.09s  0.11%  net/textproto.(*Reader).upcomingHeaderKeys
     0.03s 0.037%  0.85%      0.13s  0.16%  net/textproto.canonicalMIMEHeaderKey
     0.03s 0.037%  0.81%      0.03s 0.037%  net/url.getScheme
     0.03s 0.037%  0.78%      0.06s 0.074%  net/url.parseHost
     0.03s 0.037%  0.74%      0.03s 0.037%  runtime.(*activeSweep).begin (inline)
     0.03s 0.037%   0.7%      0.03s 0.037%  runtime.(*gcBitsArena).tryAlloc (inline)
     0.03s 0.037%  0.66%      0.03s 0.037%  runtime.(*gcWork).empty
     0.03s 0.037%  0.63%      0.03s 0.037%  runtime.(*m).clearAllpSnapshot
     0.03s 0.037%  0.59%      0.03s 0.037%  runtime.(*mspan).divideByElemSize (inline)
     0.03s 0.037%  0.55%      0.03s 0.037%  runtime.(*mspan).heapBitsSmallForAddr
     0.03s 0.037%  0.52%      0.03s 0.037%  runtime.(*mspan).refillAllocCache
     0.03s 0.037%  0.48%      0.03s 0.037%  runtime.(*stackScanState).addObject
     0.03s 0.037%  0.44%      0.03s 0.037%  runtime.(*sweepLocker).tryAcquire
     0.03s 0.037%  0.41%      0.03s 0.037%  runtime.arenaIndex (inline)
     0.03s 0.037%  0.37%      0.24s   0.3%  runtime.bulkBarrierPreWrite
     0.03s 0.037%  0.33%      0.09s  0.11%  runtime.casgstatus
     0.03s 0.037%   0.3%      0.03s 0.037%  runtime.chanrecv
     0.03s 0.037%  0.26%      0.03s 0.037%  runtime.findmoduledatap (inline)
     0.03s 0.037%  0.22%      0.18s  0.22%  runtime.gdestroy
     0.03s 0.037%  0.18%      0.03s 0.037%  runtime.getMCache (inline)
     0.03s 0.037%  0.15%      0.05s 0.062%  runtime.heapArenaOf (inline)
     0.03s 0.037%  0.11%      0.24s   0.3%  runtime.makemap
     0.03s 0.037% 0.074%      0.23s  0.28%  runtime.mapIterStart
     0.03s 0.037% 0.037%      0.03s 0.037%  runtime.mapaccess2
     0.03s 0.037%     0%      0.03s 0.037%  runtime.memequal64
     0.03s 0.037% 0.037%      0.03s 0.037%  runtime.nanotime1
     0.03s 0.037% 0.074%      0.08s 0.098%  runtime.netpoll
     0.03s 0.037%  0.11%      0.03s 0.037%  runtime.netpollblockcommit
     0.03s 0.037%  0.15%      0.27s  0.33%  runtime.pcdatavalue
     0.03s 0.037%  0.18%      0.03s 0.037%  runtime.pidlegetSpinning
     0.03s 0.037%  0.22%      0.03s 0.037%  runtime.procPin (inline)
     0.03s 0.037%  0.26%      0.13s  0.16%  runtime.ready
     0.03s 0.037%   0.3%      0.03s 0.037%  runtime.releaseSudog
     0.03s 0.037%  0.33%      0.14s  0.17%  runtime.semacquire1
     0.03s 0.037%  0.37%      0.04s 0.049%  runtime.spanSetScans
     0.03s 0.037%  0.41%      0.03s 0.037%  runtime.strequal
    -0.03s 0.037%  0.37%     -0.03s 0.037%  runtime.typePointers.nextFast (inline)
     0.03s 0.037%  0.41%      0.05s 0.062%  runtime.typehash
     0.03s 0.037%  0.44%      0.09s  0.11%  runtime.wakep
     0.03s 0.037%  0.48%      0.21s  0.26%  runtime.wbMove
     0.03s 0.037%  0.52%      0.03s 0.037%  strings.(*byteStringReplacer).Replace
     0.03s 0.037%  0.55%      0.03s 0.037%  sync.(*poolDequeue).unpack (inline)
     0.03s 0.037%  0.59%      0.03s 0.037%  sync.indexLocal (inline)
     0.03s 0.037%  0.63%      0.13s  0.16%  sync.runtime_notifyListNotifyAll
     0.03s 0.037%  0.66%      0.03s 0.037%  sync.runtime_procUnpin
     0.03s 0.037%   0.7%      0.03s 0.037%  sync/atomic.(*Bool).Load (inline)
     0.03s 0.037%  0.74%      6.11s  7.52%  syscall.write
     0.03s 0.037%  0.78%      1.27s  1.56%  time.Time.AppendFormat
     0.02s 0.025%   0.8%      0.06s 0.074%  bytes.Cut
     0.02s 0.025%  0.82%      0.04s 0.049%  bytes.Index
     0.02s 0.025%  0.85%      0.07s 0.086%  context.(*cancelCtx).Value
     0.02s 0.025%  0.87%      0.42s  0.52%  context.WithCancel.func1
    -0.02s 0.025%  0.85%     -0.40s  0.49%  database/sql.(*DB).putConn
     0.02s 0.025%  0.87%      0.04s 0.049%  fmt.(*buffer).writeString (inline)
     0.02s 0.025%   0.9%      0.09s  0.11%  fmt.(*pp).fmtString
     0.02s 0.025%  0.92%      0.04s 0.049%  fmt.(*pp).free
     0.02s 0.025%  0.95%      0.08s 0.098%  fmt.newPrinter
    -0.02s 0.025%  0.92%     -0.18s  0.22%  github.com/jackc/pgx/v5.(*Conn).Deallocate
     0.02s 0.025%  0.95%     -0.12s  0.15%  github.com/jackc/pgx/v5/pgconn.(*PgConn).Close
    -0.02s 0.025%  0.92%     -0.03s 0.037%  github.com/jackc/pgx/v5/pgconn.(*scramClient).recvServerFirstMessage
    -0.02s 0.025%   0.9%     -0.11s  0.14%  github.com/jackc/pgx/v5/pgconn.parseURLSettings
    -0.02s 0.025%  0.87%     -0.06s 0.074%  github.com/jackc/pgx/v5/pgconn/ctxwatch.(*ContextWatcher).Watch
    -0.02s 0.025%  0.85%     -0.53s  0.65%  github.com/jackc/pgx/v5/pgproto3.(*Frontend).Flush
     0.02s 0.025%  0.87%     -0.15s  0.18%  github.com/jackc/pgx/v5/stdlib.(*Conn).Close
    -0.02s 0.025%  0.85%     -0.17s  0.21%  github.com/jackc/pgx/v5/stdlib.(*Stmt).Close
     0.02s 0.025%  0.87%    -24.79s 30.52%  github.com/paulwwyvern/urlshortener/internal/handler/chihttp.(*Handler).GetURL
     0.02s 0.025%   0.9%    -26.70s 32.87%  github.com/paulwwyvern/urlshortener/internal/service/shortener.(*ShortenerService).GetURL
     0.02s 0.025%  0.92%      0.13s  0.16%  github.com/paulwwyvern/urlshortener/pkg/lrucache.(*LRUCache[go.shape.string,go.shape.string]).Get
     0.02s 0.025%  0.95%      0.08s 0.098%  go.uber.org/zap/buffer.Pool.Get
     0.02s 0.025%  0.97%      0.24s   0.3%  go.uber.org/zap/zapcore.(*jsonEncoder).AddDuration
     0.02s 0.025%     1%      0.11s  0.14%  go.uber.org/zap/zapcore.(*jsonEncoder).Clone
     0.02s 0.025%  1.02%      0.31s  0.38%  go.uber.org/zap/zapcore.EntryCaller.TrimmedPath
     0.02s 0.025%  1.05%      1.18s  1.45%  go.uber.org/zap/zapcore.ISO8601TimeEncoder
     0.02s 0.025%  1.07%      0.18s  0.22%  go.uber.org/zap/zapcore.StringDurationEncoder
     0.02s 0.025%  1.10%      0.03s 0.037%  internal/byteorder.BEUint64 (inline)
     0.02s 0.025%  1.12%      0.51s  0.63%  internal/runtime/maps.(*Map).growToSmall
     0.02s 0.025%  1.15%      0.09s  0.11%  internal/runtime/maps.(*table).reset
     0.02s 0.025%  1.17%      0.12s  0.15%  internal/runtime/maps.rand
     0.02s 0.025%  1.19%      0.05s 0.062%  internal/stringslite.IndexByte (inline)
     0.02s 0.025%  1.22%      0.12s  0.15%  internal/sync.(*Mutex).Unlock (inline)
     0.02s 0.025%  1.24%    -18.39s 22.64%  main.main.WithCompress.func7.1
     0.02s 0.025%  1.27%    -19.34s 23.81%  main.main.func1.WithAudit.1.1
    -0.02s 0.025%  1.24%     -0.30s  0.37%  net.(*sysDialer).dialParallel
     0.02s 0.025%  1.27%      1.88s  2.31%  net/http.(*chunkWriter).Write
     0.02s 0.025%  1.29%      0.08s 0.098%  net/http.(*conn).hijacked
     0.02s 0.025%  1.32%      0.05s 0.062%  net/http.(*conn).setState
     0.02s 0.025%  1.34%      0.45s  0.55%  net/http.(*connReader).abortPendingRead
     0.02s 0.025%  1.37%      1.01s  1.24%  net/http.(*connReader).backgroundRead
     0.02s 0.025%  1.39%      0.39s  0.48%  net/http.(*connReader).startBackgroundRead
     0.02s 0.025%  1.42%      0.46s  0.57%  net/http.(*response).WriteHeader
     0.02s 0.025%  1.44%      0.47s  0.58%  net/http.Header.Get (partial-inline)
     0.02s 0.025%  1.47%      0.06s 0.074%  net/http.Header.sortedKeyValues.func1
     0.02s 0.025%  1.49%    -12.59s 15.50%  net/http.serverHandler.ServeHTTP
     0.02s 0.025%  1.51%      0.08s 0.098%  net/textproto.(*Reader).ReadLine (inline)
     0.02s 0.025%  1.54%      0.08s 0.098%  net/textproto.(*Reader).readContinuedLineSlice
     0.02s 0.025%  1.56%      0.11s  0.14%  net/url.(*URL).setPath
     0.02s 0.025%  1.59%      0.51s  0.63%  net/url.parse
     0.02s 0.025%  1.61%      0.05s 0.062%  runtime.(*atomicHeadTailIndex).incTail
     0.02s 0.025%  1.64%      0.43s  0.53%  runtime.(*mcache).nextFree
     0.02s 0.025%  1.66%      0.08s 0.098%  runtime.(*mcache).releaseAll
     0.02s 0.025%  1.69%      0.29s  0.36%  runtime.(*mcentral).cacheSpan
     0.02s 0.025%  1.71%      0.06s 0.074%  runtime.(*mspan).markBitsForIndex (inline)
     0.02s 0.025%  1.74%      0.06s 0.074%  runtime.(*pageAlloc).update
     0.02s 0.025%  1.76%      0.13s  0.16%  runtime.(*sweepLocked).sweep
     0.02s 0.025%  1.79%      0.09s  0.11%  runtime.(*timer).maybeAdd
     0.02s 0.025%  1.81%      0.03s 0.037%  runtime.(*timers).adjust
     0.02s 0.025%  1.83%      0.03s 0.037%  runtime.(*unwinder).cgoCallers
     0.02s 0.025%  1.86%      0.05s 0.062%  runtime.(*unwinder).symPC
     0.02s 0.025%  1.88%      1.14s  1.40%  runtime.callers
     0.02s 0.025%  1.91%      1.12s  1.38%  runtime.callers.func1
     0.02s 0.025%  1.93%      0.08s 0.098%  runtime.concatstring3
     0.02s 0.025%  1.96%      0.20s  0.25%  runtime.entersyscall
     0.02s 0.025%  1.98%      0.11s  0.14%  runtime.funcInfo.entry (inline)
     0.02s 0.025%  2.01%      0.09s  0.11%  runtime.funcfile
     0.02s 0.025%  2.03%      0.38s  0.47%  runtime.funcline1
     0.02s 0.025%  2.06%      0.05s 0.062%  runtime.gcTrigger.test
     0.02s 0.025%  2.08%      0.03s 0.037%  runtime.getcallerfp
     0.02s 0.025%  2.11%      0.03s 0.037%  runtime.getempty
     0.02s 0.025%  2.13%      0.20s  0.25%  runtime.isSystemGoroutine
     0.02s 0.025%  2.15%      0.17s  0.21%  runtime.mallocgcSmallScanHeader
     0.02s 0.025%  2.18%      0.21s  0.26%  runtime.newInlineUnwinder
     0.02s 0.025%  2.20%      0.05s 0.062%  runtime.newMarkBits
     0.02s 0.025%  2.23%      0.49s   0.6%  runtime.newarray
     0.02s 0.025%  2.25%      0.21s  0.26%  runtime.newproc1
     0.02s 0.025%  2.28%      0.03s 0.037%  runtime.pageIndexOf (inline)
     0.02s 0.025%  2.30%      0.06s 0.074%  runtime.rawstringtmp
     0.02s 0.025%  2.33%      0.03s 0.037%  runtime.runqget (inline)
     0.02s 0.025%  2.35%      0.15s  0.18%  runtime.runqgrab
    -0.02s 0.025%  2.33%      0.09s  0.11%  runtime.stealWork
     0.02s 0.025%  2.35%      0.03s 0.037%  runtime.typedmemclr
     0.02s 0.025%  2.38%      0.04s 0.049%  sync.runtime_procPin
     0.02s 0.025%  2.40%      0.09s  0.11%  sync/atomic.(*Value).Store
     0.02s 0.025%  2.43%      7.31s  9.00%  syscall.Syscall
    -0.02s 0.025%  2.40%     -0.04s 0.049%  time.(*Timer).Reset
     0.02s 0.025%  2.43%      0.04s 0.049%  time.Duration.format
     0.02s 0.025%  2.45%      0.06s 0.074%  time.Since
     0.02s 0.025%  2.47%      0.03s 0.037%  time.absSeconds.clock (inline)
     0.01s 0.012%  2.49%      1.42s  1.75%  bufio.(*Reader).Peek
     0.01s 0.012%  2.50%      0.05s 0.062%  bufio.(*Reader).ReadLine
     0.01s 0.012%  2.51%      0.03s 0.037%  bufio.(*Reader).ReadSlice
     0.01s 0.012%  2.52%      1.41s  1.74%  bufio.(*Reader).fill
     0.01s 0.012%  2.54%      0.04s 0.049%  bytes.(*Buffer).WriteString
    -0.01s 0.012%  2.52%      0.08s 0.098%  context.(*cancelCtx).propagateCancel
     0.01s 0.012%  2.54%      0.19s  0.23%  context.WithCancel
     0.01s 0.012%  2.55%     -0.10s  0.12%  crypto/internal/fips140deps/byteorder.BEPutUint32 (inline)
     0.01s 0.012%  2.56%     -0.16s   0.2%  database/sql.(*DB).prepareDC.func1
    -0.01s 0.012%  2.55%    -15.63s 19.24%  database/sql.(*Stmt).QueryContext.func1
     0.01s 0.012%  2.56%      0.14s  0.17%  database/sql.(*driverConn).prepareLocked
     0.01s 0.012%  2.57%     -0.05s 0.062%  database/sql.withLock
     0.01s 0.012%  2.59%      0.09s  0.11%  encoding/json.stringEncoder
     0.01s 0.012%  2.60%      0.12s  0.15%  encoding/json.typeEncoder
     0.01s 0.012%  2.61%      0.13s  0.16%  fmt.(*pp).doPrint
    -0.01s 0.012%  2.60%     -0.03s 0.037%  fmt.Sprintf
     0.01s 0.012%  2.61%    -26.48s 32.60%  github.com/jackc/pgx/v5.ConnectConfig
    -0.01s 0.012%  2.60%      0.09s  0.11%  github.com/jackc/pgx/v5.ParseConfigWithOptions
     0.01s 0.012%  2.61%     -0.05s 0.062%  github.com/jackc/pgx/v5/pgconn.(*PgConn).execExtendedSuffix
    -0.01s 0.012%  2.60%     -0.58s  0.71%  github.com/jackc/pgx/v5/pgconn.(*PgConn).flushWithPotentialWriteReadDeadlock
     0.01s 0.012%  2.61%     -0.12s  0.15%  github.com/jackc/pgx/v5/pgconn.(*PgConn).peekMessage
     0.01s 0.012%  2.62%    -26.24s 32.31%  github.com/jackc/pgx/v5/pgconn.(*PgConn).scramAuth
    -0.01s 0.012%  2.61%     -0.05s 0.062%  github.com/jackc/pgx/v5/pgconn.(*scramClient).recvServerFinalMessage
     0.01s 0.012%  2.62%      0.07s 0.086%  github.com/jackc/pgx/v5/pgconn.parseEnvSettings
    -0.01s 0.012%  2.61%     -0.08s 0.098%  github.com/jackc/pgx/v5/pgconn/internal/bgreader.(*BGReader).Read
    -0.01s 0.012%  2.60%     -0.05s 0.062%  github.com/jackc/pgx/v5/pgproto3.(*ParameterStatus).Decode
    -0.01s 0.012%  2.59%     -0.08s 0.098%  github.com/jackc/pgx/v5/pgproto3.(*chunkReader).Next
     0.01s 0.012%  2.60%     -0.06s 0.074%  github.com/jackc/pgx/v5/stdlib.(*Conn).ResetSession
     0.01s 0.012%  2.61%     -0.04s 0.049%  github.com/jackc/pgx/v5/stdlib.(*Rows).Next
     0.01s 0.012%  2.62%    -26.36s 32.46%  github.com/jackc/pgx/v5/stdlib.(*driverConnector).Connect
     0.01s 0.012%  2.63%      0.47s  0.58%  github.com/paulwwyvern/urlshortener/internal/handler/middleware/logger.(*loggerResponseWriter).WriteHeader
    -0.01s 0.012%  2.62%    -26.88s 33.10%  github.com/paulwwyvern/urlshortener/internal/repository/storage/postgres.(*Storage).GetURL
     0.01s 0.012%  2.63%      0.05s 0.062%  github.com/paulwwyvern/urlshortener/pkg/httphelpers/httperr.GetError
     0.01s 0.012%  2.65%      0.06s 0.074%  go.uber.org/zap/buffer.(*Buffer).AppendInt (inline)
     0.01s 0.012%  2.66%      0.07s 0.086%  go.uber.org/zap/buffer.(*Buffer).Write (partial-inline)
     0.01s 0.012%  2.67%      0.21s  0.26%  go.uber.org/zap/internal/pool.(*Pool[go.shape.*uint8]).Get (inline)
     0.01s 0.012%  2.68%      0.05s 0.062%  go.uber.org/zap/internal/stacktrace.(*Stack).Free
     0.01s 0.012%  2.70%      0.05s 0.062%  go.uber.org/zap/zapcore.(*CheckedEntry).reset (inline)
     0.01s 0.012%  2.71%      0.26s  0.32%  go.uber.org/zap/zapcore.(*ioCore).Check
     0.01s 0.012%  2.72%      0.13s  0.16%  go.uber.org/zap/zapcore.(*jsonEncoder).AddInt64
     0.01s 0.012%  2.73%      0.67s  0.82%  go.uber.org/zap/zapcore.(*jsonEncoder).AddString
     0.01s 0.012%  2.75%      0.70s  0.86%  go.uber.org/zap/zapcore.(*jsonEncoder).safeAddString (inline)
     0.01s 0.012%  2.76%      0.38s  0.47%  go.uber.org/zap/zapcore.ShortCallerEncoder
     0.01s 0.012%  2.77%      0.03s 0.037%  go.uber.org/zap/zapcore.consoleEncoder.addSeparatorIfNecessary (inline)
     0.01s 0.012%  2.78%      0.03s 0.037%  go.uber.org/zap/zapcore.consoleEncoder.writeContext.func1
     0.01s 0.012%  2.79%      0.09s  0.11%  internal/poll.(*FD).decref
     0.01s 0.012%  2.81%      0.05s 0.062%  internal/poll.(*FD).incref (inline)
     0.01s 0.012%  2.82%      0.03s 0.037%  internal/poll.(*FD).writeUnlock
     0.01s 0.012%  2.83%      0.16s   0.2%  internal/poll.(*pollDesc).wait
     0.01s 0.012%  2.84%      0.03s 0.037%  internal/runtime/atomic.(*Bool).Load (inline)
     0.01s 0.012%  2.86%      0.21s  0.26%  internal/runtime/maps.NewMap
     0.01s 0.012%  2.87%      0.18s  0.22%  internal/strconv.AppendInt
     0.01s 0.012%  2.88%      0.04s 0.049%  internal/stringslite.Index
     0.01s 0.012%  2.89%      0.10s  0.12%  internal/sync.(*Mutex).unlockSlow
     0.01s 0.012%  2.91%      0.14s  0.17%  internal/sync.runtime_SemacquireMutex
     0.01s 0.012%  2.92%      0.09s  0.11%  internal/sync.runtime_Semrelease
     0.01s 0.012%  2.93%      0.03s 0.037%  internal/sync.runtime_nanotime
     0.01s 0.012%  2.94%     -0.07s 0.086%  io.ReadAtLeast
     0.01s 0.012%  2.95%      0.03s 0.037%  net.(*conn).Close
     0.01s 0.012%  2.97%      0.73s   0.9%  net.(*conn).SetReadDeadline
     0.01s 0.012%  2.98%      4.51s  5.55%  net.(*conn).Write
     0.01s 0.012%  2.99%      1.57s  1.93%  net.(*netFD).Read
    -0.01s 0.012%  2.98%     -0.26s  0.32%  net.(*netFD).connect
    -0.01s 0.012%  2.97%     -0.31s  0.38%  net.(*netFD).dial
     0.01s 0.012%  2.98%     -0.28s  0.34%  net.(*sysDialer).dialSingle
    -0.01s 0.012%  2.97%     -0.28s  0.34%  net.(*sysDialer).doDialTCP (inline)
     0.01s 0.012%  2.98%     -0.27s  0.33%  net.(*sysDialer).doDialTCPProto
     0.01s 0.012%  2.99%     -0.27s  0.33%  net.internetSocket
     0.01s 0.012%  3.00%      0.04s 0.049%  net/http.(*response).Write
     0.01s 0.012%  3.02%      0.03s 0.037%  net/http.(*response).write
     0.01s 0.012%  3.03%      0.07s 0.086%  net/http.fixLength
     0.01s 0.012%  3.04%      0.06s 0.074%  net/http.htmlEscape (inline)
     0.01s 0.012%  3.05%      0.05s 0.062%  net/http.requestBodyRemains
     0.01s 0.012%  3.07%      0.10s  0.12%  net/http.writeStatusLine
     0.01s 0.012%  3.08%      0.06s 0.074%  net/textproto.(*Reader).readLineSlice
     0.01s 0.012%  3.09%      0.10s  0.12%  net/textproto.MIMEHeader.Del (inline)
    -0.01s 0.012%  3.08%      0.27s  0.33%  net/textproto.MIMEHeader.Set (inline)
     0.01s 0.012%  3.09%      0.24s   0.3%  net/url.ParseRequestURI
    -0.01s 0.012%  3.08%      1.74s  2.14%  os.(*File).Write
     0.01s 0.012%  3.09%      0.03s 0.037%  runtime.(*_panic).nextDefer
     0.01s 0.012%  3.10%      0.03s 0.037%  runtime.(*gcControllerState).trigger
     0.01s 0.012%  3.11%      0.09s  0.11%  runtime.(*inlineUnwinder).next
    -0.01s 0.012%  3.10%      0.07s 0.086%  runtime.(*mcache).prepareForSweep
     0.01s 0.012%  3.11%      0.13s  0.16%  runtime.(*mcentral).grow
     0.01s 0.012%  3.13%      0.08s 0.098%  runtime.(*mheap).alloc
     0.01s 0.012%  3.14%      0.05s 0.062%  runtime.(*mheap).initSpan
     0.01s 0.012%  3.15%      0.04s 0.049%  runtime.(*mheap).nextSpanForSweep
     0.01s 0.012%  3.16%      0.17s  0.21%  runtime.(*moduledata).funcName
     0.01s 0.012%  3.18%      0.03s 0.037%  runtime.(*semaRoot).queue
     0.01s 0.012%  3.19%      0.06s 0.074%  runtime.(*spanSet).pop
     0.01s 0.012%  3.20%      0.19s  0.23%  runtime.(*timer).modify
     0.01s 0.012%  3.21%      0.03s 0.037%  runtime.(*timers).addHeap
     0.01s 0.012%  3.23%      0.05s 0.062%  runtime.(*timers).cleanHead
     0.01s 0.012%  3.24%      0.16s   0.2%  runtime.(*unwinder).initAt
     0.01s 0.012%  3.25%      0.17s  0.21%  runtime.(*wbBuf).get2 (inline)
     0.01s 0.012%  3.26%      0.22s  0.27%  runtime.adjustframe
     0.01s 0.012%  3.28%      0.04s 0.049%  runtime.atomicwb
     0.01s 0.012%  3.29%      0.15s  0.18%  runtime.bgsweep
     0.01s 0.012%  3.30%      0.03s 0.037%  runtime.cheaprandn (inline)
     0.01s 0.012%  3.31%      0.08s 0.098%  runtime.concatstring5
     0.01s 0.012%  3.32%      0.05s 0.062%  runtime.deductSweepCredit
     0.01s 0.012%  3.34%      0.08s 0.098%  runtime.forEachPInternal
    -0.01s 0.012%  3.32%     -0.04s 0.049%  runtime.futexsleep
     0.01s 0.012%  3.34%      0.57s   0.7%  runtime.gcAssistAlloc
    -0.01s 0.012%  3.32%      1.38s  1.70%  runtime.gcDrain
    -0.01s 0.012%  3.31%      0.50s  0.62%  runtime.gcDrainN
     0.01s 0.012%  3.32%      0.04s 0.049%  runtime.gcShouldScheduleWorker
     0.01s 0.012%  3.34%      0.14s  0.17%  runtime.goready (inline)
     0.01s 0.012%  3.35%      0.27s  0.33%  runtime.goschedImpl
     0.01s 0.012%  3.36%      0.46s  0.57%  runtime.heapSetTypeNoHeader (inline)
     0.01s 0.012%  3.37%     -0.03s 0.037%  runtime.injectglist
     0.01s 0.012%  3.39%      0.08s 0.098%  runtime.mapIterNext
     0.01s 0.012%  3.40%      0.04s 0.049%  runtime.mapaccess1_fast64
     0.01s 0.012%  3.41%      0.23s  0.28%  runtime.mapdelete
     0.01s 0.012%  3.42%      0.08s 0.098%  runtime.mapdelete_faststr
     0.01s 0.012%  3.44%      0.74s  0.91%  runtime.mcall
     0.01s 0.012%  3.45%      0.03s 0.037%  runtime.netpollgoready.goready.func1
     0.01s 0.012%  3.46%      0.03s 0.037%  runtime.netpollready
    -0.01s 0.012%  3.45%      0.25s  0.31%  runtime.newproc
    -0.01s 0.012%  3.44%     -0.08s 0.098%  runtime.notesleep
     0.01s 0.012%  3.45%      0.03s 0.037%  runtime.notetsleep
     0.01s 0.012%  3.46%      0.25s  0.31%  runtime.pcdatavalue1
     0.01s 0.012%  3.47%      0.07s 0.086%  runtime.profilealloc
     0.01s 0.012%  3.48%      0.13s  0.16%  runtime.readyWithTime
    -0.01s 0.012%  3.47%      0.14s  0.17%  runtime.runqsteal
     0.01s 0.012%  3.48%      0.33s  0.41%  runtime.scanframeworker
     0.01s 0.012%  3.50%      0.54s  0.66%  runtime.scanstack
     0.01s 0.012%  3.51%      0.04s 0.049%  runtime.selectnbrecv
     0.01s 0.012%  3.52%      0.05s 0.062%  runtime.signalM
     0.01s 0.012%  3.53%      0.16s   0.2%  runtime.sweepone
     0.01s 0.012%  3.55%      0.15s  0.18%  slices.pdqsortCmpFunc[go.shape.struct { net/http.key string; net/http.values []string }]
     0.01s 0.012%  3.56%      0.14s  0.17%  sync.(*Cond).Broadcast
     0.01s 0.012%  3.57%      0.14s  0.17%  sync.(*Cond).Wait
     0.01s 0.012%  3.58%      0.07s 0.086%  sync.(*poolChain).popHead
     0.01s 0.012%  3.60%      0.06s 0.074%  sync.(*poolChain).popTail
     0.01s 0.012%  3.61%      0.04s 0.049%  sync.(*poolDequeue).popTail
     0.01s 0.012%  3.62%      7.05s  8.68%  syscall.RawSyscall6
    -0.01s 0.012%  3.61%      6.10s  7.51%  syscall.Write (inline)
     0.01s 0.012%  3.62%      1.38s  1.70%  syscall.read
     0.01s 0.012%  3.63%      0.11s  0.14%  time.Duration.String (inline)
     0.01s 0.012%  3.64%      0.09s  0.11%  time.appendNano
     0.01s 0.012%  3.66%      0.07s 0.086%  time.runtimeNano
         0     0%  3.66%     -0.03s 0.037%  bytes.(*Buffer).ReadBytes (inline)
         0     0%  3.66%      0.04s 0.049%  bytes.IndexByte (inline)
         0     0%  3.66%     -0.03s 0.037%  bytes.Join
         0     0%  3.66%      0.04s 0.049%  bytes.TrimLeft
         0     0%  3.66%      0.04s 0.049%  context.(*afterFuncCtx).cancel
         0     0%  3.66%      0.41s   0.5%  context.(*cancelCtx).cancel
         0     0%  3.66%      0.16s   0.2%  context.(*valueCtx).Value
         0     0%  3.66%     -0.05s 0.062%  context.AfterFunc
         0     0%  3.66%      0.26s  0.32%  context.removeChild
         0     0%  3.66%      0.14s  0.17%  context.withCancel (inline)
         0     0%  3.66%      0.03s 0.037%  crypto/internal/fips140deps/byteorder.BEUint64 (inline)
         0     0%  3.66%    -25.72s 31.67%  crypto/pbkdf2.Key[go.shape.interface { BlockSize int; Reset; Size int; Sum []uint8; Write  }]
         0     0%  3.66%    -10.98s 13.52%  database/sql.(*DB).PrepareContext
         0     0%  3.66%    -10.98s 13.52%  database/sql.(*DB).PrepareContext.func1
         0     0%  3.66%     -0.08s 0.098%  database/sql.(*DB).addDep
         0     0%  3.66%     -0.04s 0.049%  database/sql.(*DB).addDepLocked (inline)
         0     0%  3.66%    -26.36s 32.46%  database/sql.(*DB).conn
         0     0%  3.66%    -10.98s 13.52%  database/sql.(*DB).prepare
         0     0%  3.66%     -0.07s 0.086%  database/sql.(*DB).prepareDC
         0     0%  3.66%      0.17s  0.21%  database/sql.(*DB).prepareDC.func2
         0     0%  3.66%      0.04s 0.049%  database/sql.(*DB).removeDep
         0     0%  3.66%    -26.61s 32.76%  database/sql.(*DB).retry
         0     0%  3.66%     -0.27s  0.33%  database/sql.(*Row).Scan
         0     0%  3.66%     -0.23s  0.28%  database/sql.(*Rows).Close
         0     0%  3.66%     -0.04s 0.049%  database/sql.(*Rows).Next
         0     0%  3.66%     -0.03s 0.037%  database/sql.(*Rows).Next.func1
         0     0%  3.66%     -0.24s   0.3%  database/sql.(*Rows).close
         0     0%  3.66%     -0.05s 0.062%  database/sql.(*Rows).initContextClose
         0     0%  3.66%     -0.03s 0.037%  database/sql.(*Rows).nextLocked
         0     0%  3.66%    -15.62s 19.23%  database/sql.(*Stmt).QueryContext
         0     0%  3.66%     -0.21s  0.26%  database/sql.(*Stmt).QueryContext.func1.1
         0     0%  3.66%    -15.62s 19.23%  database/sql.(*Stmt).QueryRowContext
         0     0%  3.66%    -15.49s 19.07%  database/sql.(*Stmt).connStmt
         0     0%  3.66%     -0.03s 0.037%  database/sql.(*Stmt).connStmt.func1
         0     0%  3.66%     -0.03s 0.037%  database/sql.(*Stmt).prepareOnConnLocked
         0     0%  3.66%     -0.38s  0.47%  database/sql.(*driverConn).Close
         0     0%  3.66%     -0.34s  0.42%  database/sql.(*driverConn).finalClose
         0     0%  3.66%     -0.15s  0.18%  database/sql.(*driverConn).finalClose.func2
         0     0%  3.66%     -0.40s  0.49%  database/sql.(*driverConn).releaseConn
         0     0%  3.66%     -0.06s 0.074%  database/sql.(*driverConn).resetSession
         0     0%  3.66%     -0.18s  0.22%  database/sql.(*driverStmt).Close
         0     0%  3.66%      0.14s  0.17%  database/sql.ctxDriverPrepare
         0     0%  3.66%     -0.06s 0.074%  database/sql.ctxDriverStmtQuery
         0     0%  3.66%     -0.05s 0.062%  database/sql.rowsiFromStatement
         0     0%  3.66%      0.67s  0.82%  encoding/json.(*encodeState).marshal
         0     0%  3.66%      0.09s  0.11%  encoding/json.intEncoder
         0     0%  3.66%      0.12s  0.15%  encoding/json.valueEncoder
         0     0%  3.66%      0.04s 0.049%  fmt.(*fmt).padString
         0     0%  3.66%      0.04s 0.049%  fmt.(*pp).doPrintln
         0     0%  3.66%      0.11s  0.14%  fmt.Fprintln
         0     0%  3.66%    -19.09s 23.50%  github.com/go-chi/chi/v5.(*ChainHandler).ServeHTTP
         0     0%  3.66%    -18.56s 22.85%  github.com/go-chi/chi/v5.(*Mux).routeHTTP
         0     0%  3.66%     -0.03s 0.037%  github.com/jackc/pgpassfile.ReadPassfile
         0     0%  3.66%     -0.12s  0.15%  github.com/jackc/pgx/v5.(*Conn).Close
         0     0%  3.66%      0.14s  0.17%  github.com/jackc/pgx/v5.(*Conn).Prepare
         0     0%  3.66%     -0.06s 0.074%  github.com/jackc/pgx/v5.(*Conn).Query
         0     0%  3.66%      0.09s  0.11%  github.com/jackc/pgx/v5.ParseConfig (inline)
         0     0%  3.66%    -26.47s 32.59%  github.com/jackc/pgx/v5.connect
         0     0%  3.66%      0.03s 0.037%  github.com/jackc/pgx/v5/internal/iobufpool.Get
         0     0%  3.66%      0.03s 0.037%  github.com/jackc/pgx/v5/internal/iobufpool.init.0.func1
         0     0%  3.66%      0.05s 0.062%  github.com/jackc/pgx/v5/internal/stmtcache.NewLRUCache (inline)
         0     0%  3.66%     -0.03s 0.037%  github.com/jackc/pgx/v5/pgconn.(*MultiResultReader).Close (inline)
         0     0%  3.66%     -0.03s 0.037%  github.com/jackc/pgx/v5/pgconn.(*MultiResultReader).receiveMessage
         0     0%  3.66%     -0.14s  0.17%  github.com/jackc/pgx/v5/pgconn.(*PgConn).Deallocate
         0     0%  3.66%     -0.03s 0.037%  github.com/jackc/pgx/v5/pgconn.(*PgConn).Exec
         0     0%  3.66%     -0.07s 0.086%  github.com/jackc/pgx/v5/pgconn.(*PgConn).ExecStatement
         0     0%  3.66%     -0.06s 0.074%  github.com/jackc/pgx/v5/pgconn.(*PgConn).Ping
         0     0%  3.66%      0.17s  0.21%  github.com/jackc/pgx/v5/pgconn.(*PgConn).Prepare
         0     0%  3.66%     -0.04s 0.049%  github.com/jackc/pgx/v5/pgconn.(*PgConn).enterPotentialWriteReadDeadlock (inline)
         0     0%  3.66%     -0.07s 0.086%  github.com/jackc/pgx/v5/pgconn.(*PgConn).rxSASLContinue
         0     0%  3.66%     -0.05s 0.062%  github.com/jackc/pgx/v5/pgconn.(*PgConn).rxSASLFinal
         0     0%  3.66%      0.04s 0.049%  github.com/jackc/pgx/v5/pgconn.(*ResultReader).readUntilRowDescription
         0     0%  3.66%      0.03s 0.037%  github.com/jackc/pgx/v5/pgconn.(*ResultReader).receiveMessage
         0     0%  3.66%    -25.82s 31.79%  github.com/jackc/pgx/v5/pgconn.(*scramClient).clientFinalMessage
         0     0%  3.66%      0.03s 0.037%  github.com/jackc/pgx/v5/pgconn.(*scramClient).clientFirstMessage
         0     0%  3.66%    -26.53s 32.66%  github.com/jackc/pgx/v5/pgconn.ConnectConfig
         0     0%  3.66%      0.11s  0.14%  github.com/jackc/pgx/v5/pgconn.ParseConfigWithOptions
         0     0%  3.66%      0.04s 0.049%  github.com/jackc/pgx/v5/pgconn.ParseConfigWithOptions.func1
         0     0%  3.66%     -0.06s 0.074%  github.com/jackc/pgx/v5/pgconn.computeHMAC
         0     0%  3.66%     -0.05s 0.062%  github.com/jackc/pgx/v5/pgconn.computeServerSignature
         0     0%  3.66%    -26.54s 32.68%  github.com/jackc/pgx/v5/pgconn.connectOne
         0     0%  3.66%    -26.54s 32.68%  github.com/jackc/pgx/v5/pgconn.connectPreferred
         0     0%  3.66%      0.03s 0.037%  github.com/jackc/pgx/v5/pgconn.newScramClient
         0     0%  3.66%      0.06s 0.074%  github.com/jackc/pgx/v5/pgproto3.(*AuthenticationSASL).Decode
         0     0%  3.66%      0.04s 0.049%  github.com/jackc/pgx/v5/pgproto3.NewFrontend
         0     0%  3.66%      0.04s 0.049%  github.com/jackc/pgx/v5/pgproto3.newChunkReader (inline)
         0     0%  3.66%      0.14s  0.17%  github.com/jackc/pgx/v5/stdlib.(*Conn).PrepareContext
         0     0%  3.66%     -0.06s 0.074%  github.com/jackc/pgx/v5/stdlib.(*Conn).QueryContext
         0     0%  3.66%     -0.06s 0.074%  github.com/jackc/pgx/v5/stdlib.(*Stmt).QueryContext
         0     0%  3.66%      0.04s 0.049%  github.com/paulwwyvern/urlshortener/internal/handler/middleware/logger.(*loggerResponseWriter).Write
         0     0%  3.66%      0.06s 0.074%  github.com/paulwwyvern/urlshortener/internal/handler/middleware/logger.newLoggerResponseWriter
         0     0%  3.66%      5.26s  6.48%  github.com/paulwwyvern/urlshortener/internal/service/audit.(*Publisher).Update
         0     0%  3.66%      0.03s 0.037%  github.com/paulwwyvern/urlshortener/pkg/httphelpers/httpurl.GetURL
         0     0%  3.66%      0.15s  0.18%  github.com/paulwwyvern/urlshortener/pkg/httphelpers/httpuser.GetUserID
         0     0%  3.66%      8.79s 10.82%  go.uber.org/zap.(*Logger).Info
         0     0%  3.66%      0.13s  0.16%  go.uber.org/zap/buffer.(*Buffer).Free (inline)
         0     0%  3.66%      0.05s 0.062%  go.uber.org/zap/buffer.(*Buffer).String (inline)
         0     0%  3.66%      0.13s  0.16%  go.uber.org/zap/buffer.Pool.put (inline)
         0     0%  3.66%      1.24s  1.53%  go.uber.org/zap/internal/stacktrace.Capture
         0     0%  3.66%      0.25s  0.31%  go.uber.org/zap/zapcore.(*CheckedEntry).AddCore (inline)
         0     0%  3.66%      0.06s 0.074%  go.uber.org/zap/zapcore.(*jsonEncoder).AddInt32
         0     0%  3.66%      0.19s  0.23%  go.uber.org/zap/zapcore.(*jsonEncoder).AppendDuration
         0     0%  3.66%      0.42s  0.52%  go.uber.org/zap/zapcore.(*jsonEncoder).AppendString
         0     0%  3.66%      1.51s  1.86%  go.uber.org/zap/zapcore.consoleEncoder.writeContext
         0     0%  3.66%      0.06s 0.074%  go.uber.org/zap/zapcore.putCheckedEntry (inline)
         0     0%  3.66%      0.11s  0.14%  go.uber.org/zap/zapcore.systemClock.Now
         0     0%  3.66%      0.04s 0.049%  internal/bytealg.CompareString (inline)
         0     0%  3.66%     -0.03s 0.037%  internal/bytealg.MakeNoZero
         0     0%  3.66%     -0.10s  0.12%  internal/poll.(*FD).Init
         0     0%  3.66%      0.72s  0.89%  internal/poll.(*FD).SetReadDeadline (inline)
         0     0%  3.66%      0.28s  0.34%  internal/poll.(*FD).SetWriteDeadline (inline)
         0     0%  3.66%      0.03s 0.037%  internal/poll.(*FD).writeLock (inline)
         0     0%  3.66%     -0.10s  0.12%  internal/poll.(*pollDesc).init
         0     0%  3.66%      0.16s   0.2%  internal/poll.(*pollDesc).waitRead (inline)
         0     0%  3.66%      7.51s  9.25%  internal/poll.ignoringEINTRIO (inline)
         0     0%  3.66%     -0.10s  0.12%  internal/poll.runtime_pollOpen
         0     0%  3.66%      0.15s  0.18%  internal/poll.runtime_pollWait
         0     0%  3.66%      0.49s   0.6%  internal/runtime/maps.newGroups (inline)
         0     0%  3.66%      0.09s  0.11%  internal/runtime/maps.newTable
         0     0%  3.66%      0.49s   0.6%  internal/runtime/maps.newarray
         0     0%  3.66%     -0.12s  0.15%  internal/runtime/syscall/linux.EpollCtl (inline)
         0     0%  3.66%     -0.03s 0.037%  maps.Copy[go.shape.map[string]string,go.shape.map[string]string,go.shape.string,go.shape.string] (inline)
         0     0%  3.66%     -0.30s  0.37%  net.(*Dialer).DialContext
         0     0%  3.66%      0.28s  0.34%  net.(*conn).SetWriteDeadline
         0     0%  3.66%      0.72s  0.89%  net.(*netFD).SetReadDeadline (inline)
         0     0%  3.66%      0.28s  0.34%  net.(*netFD).SetWriteDeadline (inline)
         0     0%  3.66%     -0.28s  0.34%  net.(*sysDialer).dialSerial
         0     0%  3.66%     -0.28s  0.34%  net.(*sysDialer).dialTCP
         0     0%  3.66%      0.03s 0.037%  net.setKeepAlive
         0     0%  3.66%     -0.28s  0.34%  net.socket
         0     0%  3.66%      0.03s 0.037%  net.sysSocket
         0     0%  3.66%      0.16s   0.2%  net/http.(*Request).SetPathValue (inline)
         0     0%  3.66%      0.05s 0.062%  net/http.(*chunkWriter).writeHeader.func1 (inline)
         0     0%  3.66%      0.25s  0.31%  net/http.(*conn).readRequest.func1
         0     0%  3.66%      0.36s  0.44%  net/http.Header.Clone (inline)
         0     0%  3.66%      0.10s  0.12%  net/http.Header.Del (inline)
         0     0%  3.66%      0.27s  0.33%  net/http.Header.Set (inline)
         0     0%  3.66%      0.65s   0.8%  net/http.Header.WriteSubset (inline)
         0     0%  3.66%      0.04s 0.049%  net/http.extraHeader.Write
         0     0%  3.66%      0.03s 0.037%  net/http.newBufioWriterSize
         0     0%  3.66%      0.07s 0.086%  net/http.putBufioWriter
         0     0%  3.66%      0.79s  0.97%  net/textproto.(*Reader).ReadMIMEHeader (inline)
         0     0%  3.66%      0.08s 0.098%  net/url.parseAuthority
         0     0%  3.66%      1.75s  2.15%  os.(*File).write (inline)
         0     0%  3.66%     -0.03s 0.037%  os.Open (inline)
         0     0%  3.66%     -0.03s 0.037%  os.OpenFile
         0     0%  3.66%     -0.03s 0.037%  os.open
         0     0%  3.66%     -0.03s 0.037%  os.openFileNolog
         0     0%  3.66%     -0.03s 0.037%  os.openFileNolog.func1 (inline)
         0     0%  3.66%      0.04s 0.049%  runtime.(*gcControllerState).assignWaitingGCWorker
         0     0%  3.66%      0.04s 0.049%  runtime.(*gcControllerState).findRunnableGCWorker
         0     0%  3.66%      0.25s  0.31%  runtime.(*inlineUnwinder).resolveInternal (inline)
         0     0%  3.66%      0.08s 0.098%  runtime.(*mcentral).uncacheSpan
         0     0%  3.66%      0.07s 0.086%  runtime.(*mheap).alloc.func1
         0     0%  3.66%      0.07s 0.086%  runtime.(*mheap).allocSpan
         0     0%  3.66%      0.05s 0.062%  runtime.(*mheap).freeSpan (inline)
         0     0%  3.66%      0.05s 0.062%  runtime.(*mheap).freeSpanLocked
         0     0%  3.66%      0.04s 0.049%  runtime.(*mspan).initHeapBits
         0     0%  3.66%      0.03s 0.037%  runtime.(*mspan).objIndex (inline)
         0     0%  3.66%      0.04s 0.049%  runtime.(*mspan).typePointersOfUnchecked
         0     0%  3.66%      0.03s 0.037%  runtime.(*pageAlloc).free
         0     0%  3.66%      0.19s  0.23%  runtime.(*pageAlloc).scavenge
         0     0%  3.66%      0.19s  0.23%  runtime.(*pageAlloc).scavenge.func1
         0     0%  3.66%      0.19s  0.23%  runtime.(*pageAlloc).scavengeOne
         0     0%  3.66%      0.19s  0.23%  runtime.(*scavengerState).init.func2
         0     0%  3.66%      0.19s  0.23%  runtime.(*scavengerState).run
         0     0%  3.66%      0.05s 0.062%  runtime.(*sweepLocked).sweep.(*mheap).freeSpan.func2
         0     0%  3.66%      0.70s  0.86%  runtime.(*unwinder).next
         0     0%  3.66%      1.11s  1.37%  runtime.Callers (inline)
         0     0%  3.66%      0.13s  0.16%  runtime.CallersFrames (inline)
         0     0%  3.66%      0.05s 0.062%  runtime.acquirep
         0     0%  3.66%      0.05s 0.062%  runtime.acquirepNoTrace
         0     0%  3.66%      0.19s  0.23%  runtime.bgscavenge
         0     0%  3.66%      0.37s  0.46%  runtime.copystack
         0     0%  3.66%      0.57s   0.7%  runtime.deductAssistCredit
         0     0%  3.66%      0.12s  0.15%  runtime.exitsyscallNoP
         0     0%  3.66%      0.23s  0.28%  runtime.findnull
         0     0%  3.66%      0.08s 0.098%  runtime.forEachP (inline)
         0     0%  3.66%      0.07s 0.086%  runtime.funcNameForPrint (inline)
         0     0%  3.66%      0.11s  0.14%  runtime.funcname (inline)
         0     0%  3.66%      0.41s   0.5%  runtime.funcspdelta (inline)
         0     0%  3.66%     -0.06s 0.074%  runtime.futexwakeup
         0     0%  3.66%      0.50s  0.62%  runtime.gcAssistAlloc.func2
         0     0%  3.66%      0.50s  0.62%  runtime.gcAssistAlloc1
         0     0%  3.66%      1.48s  1.82%  runtime.gcBgMarkWorker
         0     0%  3.66%      1.38s  1.70%  runtime.gcBgMarkWorker.func2
         0     0%  3.66%      1.42s  1.75%  runtime.gcDrainMarkWorkerDedicated (inline)
         0     0%  3.66%     -0.04s 0.049%  runtime.gcDrainMarkWorkerIdle (inline)
         0     0%  3.66%      0.14s  0.17%  runtime.gcMarkDone
         0     0%  3.66%      0.03s 0.037%  runtime.gcMarkDone.forEachP.func5
         0     0%  3.66%      0.03s 0.037%  runtime.gcMarkDone.func1
         0     0%  3.66%      0.07s 0.086%  runtime.gcMarkTermination
         0     0%  3.66%      0.05s 0.062%  runtime.gcMarkTermination.forEachP.func7
         0     0%  3.66%      0.03s 0.037%  runtime.gcMarkTermination.func4
         0     0%  3.66%      0.05s 0.062%  runtime.gcStart
         0     0%  3.66%      0.10s  0.12%  runtime.gcstopm
         0     0%  3.66%      0.19s  0.23%  runtime.goexit0
         0     0%  3.66%      0.03s 0.037%  runtime.goparkunlock (inline)
         0     0%  3.66%      0.18s  0.22%  runtime.gopreempt_m (inline)
         0     0%  3.66%      0.09s  0.11%  runtime.gosched_m
         0     0%  3.66%      0.23s  0.28%  runtime.gostringnocopy (inline)
         0     0%  3.66%     -0.07s 0.086%  runtime.injectglist.func1
         0     0%  3.66%      0.21s  0.26%  runtime.lock (inline)
         0     0%  3.66%      0.26s  0.32%  runtime.lockWithRank (inline)
         0     0%  3.66%     -0.08s 0.098%  runtime.mPark (inline)
         0     0%  3.66%      0.06s 0.074%  runtime.mProf_Malloc
         0     0%  3.66%      0.03s 0.037%  runtime.mProf_Malloc.func1
         0     0%  3.66%      0.63s  0.78%  runtime.markroot.func1
         0     0%  3.66%      0.12s  0.15%  runtime.markrootBlock
         0     0%  3.66%      0.03s 0.037%  runtime.memclrHasPointers
         0     0%  3.66%      0.19s  0.23%  runtime.morestack
         0     0%  3.66%      0.05s 0.062%  runtime.netpollblock
         0     0%  3.66%     -0.12s  0.15%  runtime.netpollopen
         0     0%  3.66%      0.25s  0.31%  runtime.newproc.func1
         0     0%  3.66%      0.54s  0.66%  runtime.newstack
         0     0%  3.66%     -0.07s 0.086%  runtime.notewakeup
         0     0%  3.66%      0.04s 0.049%  runtime.parkunlock_c
         0     0%  3.66%      0.05s 0.062%  runtime.preemptM (inline)
         0     0%  3.66%      0.04s 0.049%  runtime.preemptall
         0     0%  3.66%      0.03s 0.037%  runtime.preemptone
         0     0%  3.66%      0.13s  0.16%  runtime.procyield (inline)
         0     0%  3.66%     -0.04s 0.049%  runtime.rawbyteslice
         0     0%  3.66%      0.04s 0.049%  runtime.rawstring (inline)
         0     0%  3.66%      0.11s  0.14%  runtime.readyWithTime.goready.func1
         0     0%  3.66%      0.05s 0.062%  runtime.runSafePointFn
         0     0%  3.66%      0.10s  0.12%  runtime.semrelease1
         0     0%  3.66%      0.03s 0.037%  runtime.setprofilebucket
         0     0%  3.66%      0.06s 0.074%  runtime.srcFunc.name (inline)
         0     0%  3.66%     -0.09s  0.11%  runtime.startm
         0     0%  3.66%     -0.03s 0.037%  runtime.stopm
         0     0%  3.66%     -0.04s 0.049%  runtime.stringtoslicebyte
         0     0%  3.66%      0.09s  0.11%  runtime.suspendG
         0     0%  3.66%      0.18s  0.22%  runtime.sysUnused (inline)
         0     0%  3.66%      0.18s  0.22%  runtime.sysUnusedOS
         0     0%  3.66%      0.16s   0.2%  runtime.unlock (partial-inline)
         0     0%  3.66%      0.16s   0.2%  runtime.unlockWithRank (inline)
         0     0%  3.66%      0.77s  0.95%  runtime.wbBufFlush
         0     0%  3.66%      0.75s  0.92%  runtime.wbBufFlush.func1
         0     0%  3.66%      0.15s  0.18%  slices.SortFunc[go.shape.[]net/http.keyValues,go.shape.struct { net/http.key string; net/http.values []string }] (inline)
         0     0%  3.66%      0.18s  0.22%  strconv.AppendInt (inline)
         0     0%  3.66%      0.04s 0.049%  strings.Compare (inline)
         0     0%  3.66%      0.07s 0.086%  strings.Cut (inline)
         0     0%  3.66%      0.03s 0.037%  strings.LastIndex
         0     0%  3.66%      0.09s  0.11%  strings.LastIndexByte (inline)
         0     0%  3.66%      0.11s  0.14%  sync.(*Map).Load (inline)
         0     0%  3.66%      0.34s  0.42%  sync.(*Mutex).Lock (partial-inline)
         0     0%  3.66%      0.12s  0.15%  sync.(*Mutex).Unlock (partial-inline)
         0     0%  3.66%      0.09s  0.11%  sync.(*Pool).getSlow
         0     0%  3.66%      0.13s  0.16%  sync.runtime_notifyListWait
         0     0%  3.66%     -0.15s  0.18%  syscall.Connect
         0     0%  3.66%     -0.03s 0.037%  syscall.Open (inline)
         0     0%  3.66%      1.38s  1.70%  syscall.Read (inline)
         0     0%  3.66%      0.03s 0.037%  syscall.Socket
         0     0%  3.66%     -0.15s  0.18%  syscall.connect
         0     0%  3.66%     -0.03s 0.037%  syscall.openat
         0     0%  3.66%      0.03s 0.037%  syscall.socket
         0     0%  3.66%      0.98s  1.21%  time.Time.Format
         0     0%  3.66%      0.04s 0.049%  time.Time.Sub
```