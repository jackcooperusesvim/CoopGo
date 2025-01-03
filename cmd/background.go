package main

import "github.com/jackcooperusesvim/coopGo/model/auth"

q, ctx, err := auth.DbInfo()
return q.PubliclyUnaliveTokens(ctx)
