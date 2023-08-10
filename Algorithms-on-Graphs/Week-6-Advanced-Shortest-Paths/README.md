# Advanced Shortest Paths (Optinal)
## Bidirectional Dijkstra

> [!NOTE]
> ```
> BidirectionalDijkstra(G,s,t)
> G_R ← ReverseGraph(G)
> Filldist,dist_R with +∞ for each node
> dist[s] ← 0,dist_R[t] ← 0
> Fill prev,prev_R with None foreach node 
> proc ← empty,proc_R ← empty 
> do:
>     v ← ExtractMin(dist)
>     Process(v,G,dist,prev,proc) 
>     if v in proc_R:
>         return ShortestPath(s,dist,prev,proc,t,...)
>     v_R ← ExtractMin(dist_R)
>     repeat symmetrically for v_R as for v 
> while True
> ```
> ```
> Relax(u,v,dist,prev) 
> if dist[v] > dist[u] + w(u,v):
>     dist[v] ← dist[u] + w(u,v)
>     prev[v] ← u
> ```
> ```
> Process(u,G,dist,prev,proc) 
> for(u,v) ∈ E(G): 
>     Relax(u,v,dist,prev) 
> proc.Append(u)
> ```
> ```
> ShortestPath(s,dist,prev,proc,t,dist_R,prev_R,proc_R)
> distance ← +∞,u_best ← None
> for u in proc + proc_R:
>     if dist[u] + dist_R[u] < distance:
>         u_best ← u
>         distance ← dist[u] + dist_R[u]
> path ← empty
> last ← u_best 
> while last != s: 
>     path.Append(last)
>     last ← prev[last]
> path ← Reverse(path)
> last ← u_best 
> while last != t: 
>     last ← prev_R[last]
>     path.Append(last) 
> return(distance,path)
> ```
