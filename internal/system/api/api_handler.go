package api

import (
	"github.com/GoSimplicity/AI-CloudOps/internal/model"
	"github.com/GoSimplicity/AI-CloudOps/internal/system/service"
	"github.com/GoSimplicity/AI-CloudOps/pkg/utils/apiresponse"
	ijwt "github.com/GoSimplicity/AI-CloudOps/pkg/utils/jwt"
	"github.com/gin-gonic/gin"
)

/*
 * MIT License
 *
 * Copyright (c) 2024 Bamboo
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 *
 */

type AuthApiHandler struct {
	apiService service.AuthApiService
}

func NewAuthApiHandler(apiService service.AuthApiService) *AuthApiHandler {
	return &AuthApiHandler{
		apiService: apiService,
	}
}

func (a *AuthApiHandler) RegisterRouters(server *gin.Engine) {
	authGroup := server.Group("/api/auth")

	// API 管理相关路由
	authGroup.GET("/api/list", a.GetApiList)
	authGroup.GET("/api/all", a.GetApiListAll)
	authGroup.DELETE("/api/:id", a.DeleteApi)
	authGroup.POST("/api/create", a.CreateApi)
	authGroup.POST("/api/update", a.UpdateApi)
}

func (a *AuthApiHandler) GetApiList(ctx *gin.Context) {
	uc := ctx.MustGet("user").(ijwt.UserClaims)

	Apis, err := a.apiService.GetApiList(ctx, uc.Uid)
	if err != nil {
		apiresponse.ErrorWithMessage(ctx, err.Error())
		return
	}

	apiresponse.SuccessWithData(ctx, Apis)
}

func (a *AuthApiHandler) GetApiListAll(ctx *gin.Context) {
	Apis, err := a.apiService.GetApiListAll(ctx)
	if err != nil {
		apiresponse.ErrorWithMessage(ctx, err.Error())
		return
	}

	apiresponse.SuccessWithData(ctx, Apis)
}

func (a *AuthApiHandler) DeleteApi(ctx *gin.Context) {
	id := ctx.Param("id")

	err := a.apiService.DeleteApi(ctx, id)
	if err != nil {
		apiresponse.ErrorWithMessage(ctx, err.Error())
		return
	}

	apiresponse.SuccessWithMessage(ctx, "删除成功")
}

func (a *AuthApiHandler) CreateApi(ctx *gin.Context) {
	var ma model.Api

	if err := ctx.ShouldBindJSON(&ma); err != nil {
		apiresponse.ErrorWithMessage(ctx, err.Error())
		return
	}

	err := a.apiService.CreateApi(ctx, &ma)
	if err != nil {
		apiresponse.ErrorWithMessage(ctx, err.Error())
		return
	}

	apiresponse.SuccessWithMessage(ctx, "创建成功")
}

func (a *AuthApiHandler) UpdateApi(ctx *gin.Context) {
	var ma model.Api

	if err := ctx.ShouldBindJSON(&ma); err != nil {
		apiresponse.ErrorWithMessage(ctx, err.Error())
		return
	}

	err := a.apiService.UpdateApi(ctx, &ma)
	if err != nil {
		apiresponse.ErrorWithMessage(ctx, err.Error())
		return
	}

	apiresponse.SuccessWithMessage(ctx, "更新成功")
}
