// Copyright 2020 Google LLC
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

package fake

import (
	"context"
	"fmt"
	"log"
	"os"

	translate "cloud.google.com/go/translate/apiv3"
	"cloud.google.com/go/translate/apiv3/translatepb"
)

// TranslateTextWithConcreteClient translates text to the targetLang using the
// provided client.
func TranslateTextWithConcreteClient(client *translate.TranslationClient, text string, targetLang string) (string, error) {
	ctx := context.Background()
	log.Printf("Translating %q to %q", text, targetLang)
	req := &translatepb.TranslateTextRequest{
		Parent:             fmt.Sprintf("projects/%s/locations/global", os.Getenv("GOOGLE_CLOUD_PROJECT")),
		TargetLanguageCode: "en-US",
		Contents:           []string{text},
	}
	resp, err := client.TranslateText(ctx, req)
	if err != nil {
		return "", fmt.Errorf("unable to translate text: %v", err)
	}
	translations := resp.GetTranslations()
	if len(translations) != 1 {
		return "", fmt.Errorf("expected only one result, got %d", len(translations))
	}
	return translations[0].TranslatedText, nil
}
