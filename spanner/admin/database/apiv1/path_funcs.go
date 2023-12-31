// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package database

// DatabaseAdminInstancePath returns the path for the instance resource.
//
// Deprecated: Use
//
//	fmt.Sprintf("projects/%s/instances/%s", project, instance)
//
// instead.
func DatabaseAdminInstancePath(project, instance string) string {
	return "" +
		"projects/" +
		project +
		"/instances/" +
		instance +
		""
}

// DatabaseAdminDatabasePath returns the path for the database resource.
//
// Deprecated: Use
//
//	fmt.Sprintf("projects/%s/instances/%s/databases/%s", project, instance, database)
//
// instead.
func DatabaseAdminDatabasePath(project, instance, database string) string {
	return "" +
		"projects/" +
		project +
		"/instances/" +
		instance +
		"/databases/" +
		database +
		""
}
