/*
 * Copyright (c) 2026 The XGo Authors (xgo.dev). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package tspath

import (
	"github.com/microsoft/typescript-go/internal/tspath"
)

type Path = tspath.Path

func GetNormalizedAbsolutePath(fileName string, currentDirectory string) string {
	return tspath.GetNormalizedAbsolutePath(fileName, currentDirectory)
}

func ToPath(fileName string, basePath string, useCaseSensitiveFileNames bool) Path {
	return tspath.ToPath(fileName, basePath, useCaseSensitiveFileNames)
}
