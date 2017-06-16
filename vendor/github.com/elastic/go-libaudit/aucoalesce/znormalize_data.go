// mknormalize_data.go
// MACHINE GENERATED BY THE ABOVE COMMAND; DO NOT EDIT

// Copyright 2017 Elasticsearch Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aucoalesce

import (
	"encoding/base64"
	"fmt"
)

var assets map[string][]byte

func asset(key string) ([]byte, error) {
	if assets == nil {
		assets = map[string][]byte{}

		var value []byte
		value, _ = base64.StdEncoding.DecodeString("LS0tCiMgTWFjcm9zIGRlY2xhcmVzIHNvbWUgWUFNTCBhbmNob3JzIHRoYXQgY2FuIGJlIHJlZmVyZW5jZWQgZm9yIHNvbWUgY29tbW9uCiMgb2JqZWN0IHR5cGUgbm9ybWFsaXphdGlvbnMgbGlrZSB1c2VyLXNlc3Npb24sIHNvY2tldCwgb3IgcHJvY2Vzcy4KbWFjcm9zOgotICZkZWZhdWx0cwogIHN1YmplY3Q6CiAgICBwcmltYXJ5OiBhdWlkCiAgICBzZWNvbmRhcnk6IHVpZAogIGhvdzogW2V4ZSwgY29tbV0KCi0gJm1hY3JvLXVzZXItc2Vzc2lvbgogIHN1YmplY3Q6CiAgICBwcmltYXJ5OiBhdWlkCiAgICBzZWNvbmRhcnk6IFthY2N0LCBpZCwgdWlkXQogIG9iamVjdDoKICAgIHByaW1hcnk6IHRlcm1pbmFsCiAgICBzZWNvbmRhcnk6IFthZGRyLCBob3N0bmFtZV0KICAgIHdoYXQ6IHVzZXItc2Vzc2lvbgogIGhvdzogW2V4ZSwgdGVybWluYWxdCgotICZtYWNyby1zb2NrZXQKICA8PDogKmRlZmF1bHRzCiAgb2JqZWN0OgogICAgcHJpbWFyeTogW2FkZHIsIHBhdGhdCiAgICBzZWNvbmRhcnk6IHBvcnQKICAgIHdoYXQ6IHNvY2tldAoKLSAmbWFjcm8tcHJvY2VzcwogIDw8OiAqZGVmYXVsdHMKICBvYmplY3Q6CiAgICBwcmltYXJ5OiBbY21kLCBleGUsIGNvbW1dCiAgICBzZWNvbmRhcnk6IHBpZAogICAgd2hhdDogcHJvY2VzcwogIGhvdzogdGVybWluYWwKCiMgTm9ybWFsaXphdGlvbnMgaXMgYSBsaXN0IG9mIGRlY2xhcmF0aW9ucyBzcGVjaWZ5aW5nIGhvdyB0byBub3JtYWxpemUgdGhlIGRhdGEKIyBjb250YWluZWQgaW4gYW4gZXZlbnQuIFRoZSBub3JtYWxpemF0aW9uIGNhbiBiZSBhcHBsaWVkIGJhc2VkIG9uIHRoZSBzeXNjYWxsCiMgbmFtZSAoZS5nLiBjb25uZWN0LCBvcGVuKSBvciBiYXNlZCBvbiB0aGUgcmVjb3JkIHR5cGUgKGUuZy4gVVNFUl9MT0dJTikuCiMgTm8gdHdvIG5vcm1hbGl6YXRpb25zIGNhbiBhcHBseSB0byB0aGUgc2FtZSBzeXNjYWxsIG9yIHJlY29yZCB0eXBlLiBUaGlzCiMgd2lsbCByZXN1bHQgaW4gYSBmYWlsdXJlIGF0IGxvYWQgdGltZS4KIwojIEVhY2ggbm9ybWFsaXphdGlvbiBzaG91bGQgc3BlY2lmeToKIyAgIGFjdGlvbiAtIHdoYXQgaGFwcGVuZWQKIyAgIGFjdG9yICAtIHdobyBkaWQgdGhpcyBvciB3aG8gdHJpZ2dlcmVkIHRoZSBldmVudAojICAgb2JqZWN0IC0gd2hhdCB3YXMgdGhlICJ0aGluZyIgaW52b2x2ZWQgaW4gdGhlIGFjdGlvbiAoZS5nLiBwcm9jZXNzLCBzb2NrZXQpCiMgICBob3cgICAgLSBob3cgd2FzIHRoZSBhY3Rpb24gcGVyZm9ybWVkIChlLmcuIGV4ZSBvciB0ZXJtaW5hbCkKbm9ybWFsaXphdGlvbnM6Ci0KICBhY3Rpb246IG9wZW5lZC1maWxlCiAgb2JqZWN0OgogICAgd2hhdDogZmlsZQogIHN5c2NhbGxzOgogIC0gY3JlYXQKICAtIGZhbGxvY2F0ZQogIC0gdHJ1bmNhdGUKICAtIGZ0cnVuY2F0ZQogIC0gb3BlbgogIC0gb3BlbmF0CiAgLSByZWFkbGluawogIC0gcmVhZGxpbmthdAotCiAgYWN0aW9uOiBjaGFuZ2VkLWZpbGUtYXR0cmlidXRlcy1vZgogIG9iamVjdDoKICAgIHdoYXQ6IGZpbGUKICBzeXNjYWxsczoKICAtIHNldHhhdHRyCiAgLSBmc2V0eGF0dHIKICAtIGxzZXR4YXR0cgogIC0gcmVtb3ZleGF0dHIKICAtIGZyZW1vdmV4YXR0cgogIC0gbHJlbW92ZXhhdHRyCi0KICBhY3Rpb246IGNoYW5nZWQtZmlsZS1wZXJtaXNzaW9ucy1vZgogIG9iamVjdDoKICAgIHdoYXQ6IGZpbGUKICBzeXNjYWxsczoKICAtIGNobW9kCiAgLSBmY2htb2QKICAtIGZjaG1vZGF0Ci0KICBhY3Rpb246IGNoYW5nZWQtZmlsZS1vd25lcnNoaXAtb2YKICBvYmplY3Q6CiAgICB3aGF0OiBmaWxlCiAgc3lzY2FsbHM6CiAgLSBjaG93bgogIC0gZmNob3duCiAgLSBmY2hvd25hdAogIC0gbGNob3duCi0KICBhY3Rpb246IGxvYWRlZC1rZXJuZWwtbW9kdWxlCiAgb2JqZWN0OgogICAgd2hhdDogZmlsZQogICAgcHJpbWFyeTogbmFtZQogIHJlY29yZF90eXBlczoKICAtIEtFUk5fTU9EVUxFCiAgc3lzY2FsbHM6CiAgLSBmaW5pdF9tb2R1bGUKICAtIGluaXRfbW9kdWxlCi0KICBhY3Rpb246IHVubG9hZGVkLWtlcm5lbC1tb2R1bGUKICBvYmplY3Q6CiAgICB3aGF0OiBmaWxlCiAgc3lzY2FsbHM6CiAgLSBkZWxldGVfbW9kdWxlCi0KICBhY3Rpb246IGNyZWF0ZWQtZGlyZWN0b3J5CiAgb2JqZWN0OgogICAgd2hhdDogZmlsZQogICAgcGF0aF9pbmRleDogMQogIHN5c2NhbGxzOgogIC0gbWtkaXIKICAtIG1rZGlyYXQKLQogIGFjdGlvbjogbW91bnRlZAogIG9iamVjdDoKICAgIHdoYXQ6IGZpbGVzeXN0ZW0KICAgIHBhdGhfaW5kZXg6IDEKICBzeXNjYWxsczoKICAtIG1vdW50Ci0KICBhY3Rpb246IHJlbmFtZWQKICBvYmplY3Q6CiAgICB3aGF0OiBmaWxlCiAgICBwYXRoX2luZGV4OiAyCiAgc3lzY2FsbHM6CiAgLSByZW5hbWUKICAtIHJlbmFtZWF0CiAgLSByZW5hbWVhdDIKLQogIGFjdGlvbjogY2hlY2tlZC1tZXRhZGF0YS1vZgogIG9iamVjdDoKICAgIHdoYXQ6IGZpbGUKICBzeXNjYWxsczoKICAtIGFjY2VzcwogIC0gZmFjY2Vzc2F0CiAgLSBuZXdmc3RhdGF0CiAgLSBzdGF0CiAgLSBmc3RhdAogIC0gbHN0YXQKICAtIHN0YXQ2NAogIC0gZ2V0eGF0dHIKICAtIGxnZXR4YXR0cgogIC0gZmdldHhhdHRyCi0KICBhY3Rpb246IGNoZWNrZWQtZmlsZXN5c3RlbS1tZXRhZGF0YS1vZgogIG9iamVjdDoKICAgIHdoYXQ6IGZpbGVzeXN0ZW0KICBzeXNjYWxsczoKICAtIHN0YXRmcwogIC0gZnN0YXRmcwotCiAgYWN0aW9uOiBzeW1saW5rZWQKICBvYmplY3Q6CiAgICB3aGF0OiBmaWxlCiAgc3lzY2FsbHM6CiAgLSBzeW1saW5rCiAgLSBzeW1saW5rYXQKLQogIGFjdGlvbjogdW5tb3VudGVkCiAgb2JqZWN0OgogICAgd2hhdDogZmlsZXN5c3RlbQogIHN5c2NhbGxzOgogIC0gdW1vdW50MgotCiAgYWN0aW9uOiBkZWxldGVkCiAgb2JqZWN0OgogICAgd2hhdDogZmlsZQogIHN5c2NhbGxzOgogIC0gcm1kaXIKICAtIHVubGluawogIC0gdW5saW5rYXQKLQogIGFjdGlvbjogY2hhbmdlZC10aW1lc3RhbXAtb2YKICBvYmplY3Q6CiAgICB3aGF0OiBmaWxlCiAgc3lzY2FsbHM6CiAgLSB1dGltZQogIC0gdXRpbWVzCiAgLSBmdXRpbWVzYXQKICAtIGZ1dGltZW5zCiAgLSB1dGltZW5zYXQKLQogIGFjdGlvbjogZXhlY3V0ZWQKICBvYmplY3Q6CiAgICB3aGF0OiBmaWxlCiAgc3lzY2FsbHM6CiAgLSBleGVjdmUKICAtIGV4ZWN2ZWF0Ci0KICBhY3Rpb246IGxpc3Rlbi1mb3ItY29ubmVjdGlvbnMKICBvYmplY3Q6CiAgICB3aGF0OiBzb2NrZXQKICBzeXNjYWxsczoKICAtIGxpc3RlbgotCiAgYWN0aW9uOiBhY2NlcHRlZC1jb25uZWN0aW9uLWZyb20KICBvYmplY3Q6CiAgICB3aGF0OiBzb2NrZXQKICBzeXNjYWxsczoKICAtIGFjY2VwdAogIC0gYWNjZXB0NAotCiAgYWN0aW9uOiBib3VuZC1zb2NrZXQKICBvYmplY3Q6CiAgICB3aGF0OiBzb2NrZXQKICBzeXNjYWxsczoKICAtIGJpbmQKLQogIGFjdGlvbjogY29ubmVjdGVkLXRvCiAgb2JqZWN0OgogICAgd2hhdDogc29ja2V0CiAgc3lzY2FsbHM6CiAgLSBjb25uZWN0Ci0KICBhY3Rpb246IHJlY2VpdmVkLWZyb20KICBvYmplY3Q6CiAgICB3aGF0OiBzb2NrZXQKICBzeXNjYWxsczoKICAtIHJlY3Zmcm9tCiAgLSByZWN2bXNnCi0KICBhY3Rpb246IHNlbnQtdG8KICBvYmplY3Q6CiAgICB3aGF0OiBzb2NrZXQKICBzeXNjYWxsczoKICAtIHNlbmR0bwogIC0gc2VuZG1zZwotCiAgYWN0aW9uOiBraWxsZWQtcGlkCiAgb2JqZWN0OgogICAgd2hhdDogcHJvY2VzcwogIHN5c2NhbGxzOgogIC0ga2lsbAogIC0gdGtpbGwKICAtIHRna2lsbAotCiAgYWN0aW9uOiBjaGFuZ2VkLWlkZW50aXR5LW9mCiAgb2JqZWN0OgogICAgd2hhdDogcHJvY2VzcwogIGhvdzogc3lzY2FsbAogIHN5c2NhbGxzOgogIC0gc2V0dWlkCiAgLSBzZXRldWlkCiAgLSBzZXRmc3VpZAogIC0gc2V0cmV1aWQKICAtIHNldHJlc3VpZAogIC0gc2V0Z2lkCiAgLSBzZXRlZ2lkCiAgLSBzZXRmc2dpZAogIC0gc2V0cmVnaWQKICAtIHNldHJlc2dpZAotCiAgYWN0aW9uOiBjaGFuZ2VkLXN5c3RlbS10aW1lCiAgb2JqZWN0OgogICAgd2hhdDogc3lzdGVtCiAgc3lzY2FsbHM6CiAgLSBzZXR0aW1lb2ZkYXkKICAtIGNsb2NrX3NldHRpbWUKICAtIHN0aW1lCiAgLSBhZGp0aW1leAotCiAgYWN0aW9uOiBtYWtlLWRldmljZQogIG9iamVjdDoKICAgIHdoYXQ6IGZpbGUKICBzeXNjYWxsczoKICAtIG1rbm9kCiAgLSBta25vZGF0Ci0KICBhY3Rpb246IGNoYW5nZWQtc3lzdGVtLW5hbWUKICBvYmplY3Q6CiAgICB3aGF0OiBzeXN0ZW0KICBzeXNjYWxsczoKICAtIHNldGhvc3RuYW1lCiAgLSBzZXRkb21haW5uYW1lCi0KICBhY3Rpb246IGFsbG9jYXRlZC1tZW1vcnkKICBvYmplY3Q6CiAgICB3aGF0OiBtZW1vcnkKICBzeXNjYWxsczoKICAtIG1tYXAKICAtIGJyawotCiAgYWN0aW9uOiBhZGp1c3RlZC1zY2hlZHVsaW5nLXBvbGljeS1vZgogIG9iamVjdDoKICAgIHdoYXQ6IHByb2Nlc3MKICBob3c6IHN5c2NhbGwKICBzeXNjYWxsczoKICAtIHNjaGVkX3NldHBhcmFtCiAgLSBzY2hlZF9zZXRzY2hlZHVsZXIKICAtIHNjaGVkX3NldGF0dHIKLQogIGFjdGlvbjogY2F1c2VkLW1hYy1wb2xpY3ktZXJyb3IKICBvYmplY3Q6CiAgICB3aGF0OiBzeXN0ZW0KICByZWNvcmRfdHlwZXM6IFNFTElOVVhfRVJSCi0KICBhY3Rpb246IGxvYWRlZC1maXJld2FsbC1ydWxlLXRvCiAgb2JqZWN0OgogICAgcHJpbWFyeTogdGFibGUKICAgIHdoYXQ6IGZpcmV3YWxsCiAgcmVjb3JkX3R5cGVzOiBORVRGSUxURVJfQ0ZHCi0KICAjIENvdWxkIGJlIGVudGVyZWQgb3IgZXhpdGVkIGJhc2VkIG9uIHByb20gZmllbGQuCiAgYWN0aW9uOiBjaGFuZ2VkLXByb21pc2N1b3VzLW1vZGUtb24tZGV2aWNlCiAgb2JqZWN0OgogICAgcHJpbWFyeTogZGV2CiAgICB3aGF0OiBuZXR3b3JrLWRldmljZQogIHJlY29yZF90eXBlczogQU5PTV9QUk9NSVNDVU9VUwotCiAgYWN0aW9uOiBsb2NrZWQtYWNjb3VudAogIHJlY29yZF90eXBlczogQUNDVF9MT0NLCi0KICBhY3Rpb246IHVubG9ja2VkLWFjY291bnQKICByZWNvcmRfdHlwZXM6IEFDQ1RfVU5MT0NLCi0KICBhY3Rpb246IGFkZGVkLWdyb3VwLWFjY291bnQtdG8KICBvYmplY3Q6CiAgICBwcmltYXJ5OiBbaWQsIGFjY3RdCiAgICB3aGF0OiBhY2NvdW50CiAgcmVjb3JkX3R5cGVzOiBBRERfR1JPVVAKLQogIGFjdGlvbjogYWRkZWQtdXNlci1hY2NvdW50CiAgb2JqZWN0OgogICAgcHJpbWFyeTogW2lkLCBhY2N0XQogICAgd2hhdDogYWNjb3VudAogIHJlY29yZF90eXBlczogQUREX1VTRVIKLQogIGFjdGlvbjogY3Jhc2hlZC1wcm9ncmFtCiAgb2JqZWN0OgogICAgcHJpbWFyeTogW2NvbW0sIGV4ZV0KICAgIHNlY29uZGFyeTogcGlkCiAgICB3aGF0OiBwcm9jZXNzCiAgaG93OiBzaWcKICByZWNvcmRfdHlwZXM6IEFOT01fQUJFTkQKLQogIGFjdGlvbjogYXR0ZW1wdGVkLWV4ZWN1dGlvbi1vZi1mb3JiaWRkZW4tcHJvZ3JhbQogIG9iamVjdDoKICAgIHByaW1hcnk6IGNtZAogICAgd2hhdDogcHJvY2VzcwogIGhvdzogdGVybWluYWwKICByZWNvcmRfdHlwZXM6IEFOT01fRVhFQwotCiAgYWN0aW9uOiB1c2VkLXN1c3BjaW91cy1saW5rCiAgcmVjb3JkX3R5cGVzOiBBTk9NX0xJTksKLQogIDw8OiAqbWFjcm8tdXNlci1zZXNzaW9uCiAgYWN0aW9uOiBmYWlsZWQtbG9nLWluLXRvby1tYW55LXRpbWVzLXRvCiAgcmVjb3JkX3R5cGVzOiBBTk9NX0xPR0lOX0ZBSUxVUkVTCi0KICA8PDogKm1hY3JvLXVzZXItc2Vzc2lvbgogIGFjdGlvbjogYXR0ZW1wdGVkLWxvZy1pbi1mcm9tLXVudXN1YWwtcGxhY2UtdG8KICByZWNvcmRfdHlwZXM6IEFOT01fTE9HSU5fTE9DQVRJT04KLQogIDw8OiAqbWFjcm8tdXNlci1zZXNzaW9uCiAgYWN0aW9uOiBvcGVuZWQtdG9vLW1hbnktc2Vzc2lvbnMtdG8KICByZWNvcmRfdHlwZXM6IEFOT01fTE9HSU5fU0VTU0lPTlMKLQogIDw8OiAqbWFjcm8tdXNlci1zZXNzaW9uCiAgYWN0aW9uOiBhdHRlbXB0ZWQtbG9nLWluLWR1cmluZy11bnVzdWFsLWhvdXItdG8KICByZWNvcmRfdHlwZXM6IEFOT01fTE9HSU5fVElNRQotCiAgYWN0aW9uOiB0ZXN0ZWQtZmlsZS1zeXN0ZW0taW50ZWdyaXR5LW9mCiAgb2JqZWN0OgogICAgcHJpbWFyeTogaG9zdG5hbWUKICAgIHdoYXQ6IGZpbGVzeXN0ZW0KICByZWNvcmRfdHlwZXM6IEFOT01fUkJBQ19JTlRFR1JJVFlfRkFJTAotCiAgYWN0aW9uOiB2aW9sYXRlZC1zZWxpbnV4LXBvbGljeQogIHN1YmplY3Q6CiAgICBwcmltYXJ5OiBzY29udGV4dAogIG9iamVjdDoKICAgIHByaW1hcnk6IHRjb250ZXh0CiAgcmVjb3JkX3R5cGVzOiBBVkMKLQogIGFjdGlvbjogY2hhbmdlZC1ncm91cAogIHJlY29yZF90eXBlczogQ0hHUlBfSUQKLQogIGFjdGlvbjogY2hhbmdlZC11c2VyLWlkCiAgcmVjb3JkX3R5cGVzOiBDSFVTRVJfSUQKLQogIGFjdGlvbjogY2hhbmdlZC1hdWRpdC1jb25maWd1cmF0aW9uCiAgb2JqZWN0OgogICAgcHJpbWFyeTogW29wLCBrZXksIGF1ZGl0X2VuYWJsZWQsIGF1ZGl0X3BpZCwgYXVkaXRfYmFja2xvZ19saW1pdCwgYXVkaXRfZmFpbHVyZV0KICAgIHdoYXQ6IGF1ZGl0LWNvbmZpZwogIHJlY29yZF90eXBlczogQ09ORklHX0NIQU5HRQotCiAgPDw6ICptYWNyby11c2VyLXNlc3Npb24KICBhY3Rpb246IGFjcXVpcmVkLWNyZWRlbnRpYWxzCiAgcmVjb3JkX3R5cGVzOiBDUkVEX0FDUQotCiAgPDw6ICptYWNyby11c2VyLXNlc3Npb24KICBhY3Rpb246IGRpc3Bvc2VkLWNyZWRlbnRpYWxzCiAgcmVjb3JkX3R5cGVzOiBDUkVEX0RJU1AKLQogIDw8OiAqbWFjcm8tdXNlci1zZXNzaW9uCiAgYWN0aW9uOiByZWZyZXNoZWQtY3JlZGVudGlhbHMKICByZWNvcmRfdHlwZXM6IENSRURfUkVGUgotCiAgPDw6ICptYWNyby11c2VyLXNlc3Npb24KICBhY3Rpb246IG5lZ290aWF0ZWQtY3J5cHRvLWtleQogIG9iamVjdDoKICAgIHByaW1hcnk6IGZwCiAgICBzZWNvbmRhcnk6IFthZGRyLCBob3N0bmFtZV0KICAgIHdoYXQ6IHVzZXItc2Vzc2lvbgogIHJlY29yZF90eXBlczogQ1JZUFRPX0tFWV9VU0VSCi0KICBhY3Rpb246IGNyeXB0by1vZmZpY2VyLWxvZ2dlZC1pbgogIHJlY29yZF90eXBlczogQ1JZUFRPX0xPR0lOCi0KICBhY3Rpb246IGNyeXB0by1vZmZpY2VyLWxvZ2dlZC1vdXQKICByZWNvcmRfdHlwZXM6IENSWVBUT19MT0dPVVQKLQogIDw8OiAqbWFjcm8tdXNlci1zZXNzaW9uCiAgYWN0aW9uOiBzdGFydGVkLWNyeXB0by1zZXNzaW9uCiAgb2JqZWN0OgogICAgcHJpbWFyeTogYWRkcgogICAgc2Vjb25kYXJ5OiBbcnBvcnRdCiAgcmVjb3JkX3R5cGVzOiBDUllQVE9fU0VTU0lPTgotCiAgYWN0aW9uOiBhY2Nlc3MtcmVzdWx0CiAgcmVjb3JkX3R5cGVzOiBEQUNfQ0hFQ0sKLQogIGFjdGlvbjogYWJvcnRlZC1hdWRpdGQtc3RhcnR1cAogIG9iamVjdDoKICAgIHdoYXQ6IHNlcnZpY2UKICByZWNvcmRfdHlwZXM6IERBRU1PTl9BQk9SVAotCiAgYWN0aW9uOiByZW1vdGUtYXVkaXQtY29ubmVjdGVkCiAgb2JqZWN0OgogICAgd2hhdDogc2VydmljZQogIHJlY29yZF90eXBlczogREFFTU9OX0FDQ0VQVAotCiAgYWN0aW9uOiByZW1vdGUtYXVkaXQtZGlzY29ubmVjdGVkCiAgb2JqZWN0OgogICAgd2hhdDogc2VydmljZQogIHJlY29yZF90eXBlczogREFFTU9OX0NMT1NFCi0KICBhY3Rpb246IGNoYW5nZWQtYXVkaXRkLWNvbmZpZ3VyYXRpb24KICBvYmplY3Q6CiAgICB3aGF0OiBzZXJ2aWNlCiAgcmVjb3JkX3R5cGVzOiBEQUVNT05fQ09ORklHCi0KICBhY3Rpb246IHNodXRkb3duLWF1ZGl0CiAgb2JqZWN0OgogICAgd2hhdDogc2VydmljZQogIHJlY29yZF90eXBlczogREFFTU9OX0VORAotCiAgYWN0aW9uOiBhdWRpdC1lcnJvcgogIG9iamVjdDoKICAgIHdoYXQ6IHNlcnZpY2UKICByZWNvcmRfdHlwZXM6IERBRU1PTl9FUlIKLQogIGFjdGlvbjogcmVjb25maWd1cmVkLWF1ZGl0ZAogIG9iamVjdDoKICAgIHdoYXQ6IHNlcnZpY2UKICByZWNvcmRfdHlwZXM6IERBRU1PTl9SRUNPTkZJRwotCiAgYWN0aW9uOiByZXN1bWVkLWF1ZGl0LWxvZ2dpbmcKICBvYmplY3Q6CiAgICB3aGF0OiBzZXJ2aWNlCiAgcmVjb3JkX3R5cGVzOiBEQUVNT05fUkVTVU1FCi0KICBhY3Rpb246IHJvdGF0ZWQtYXVkaXQtbG9ncwogIG9iamVjdDoKICAgIHdoYXQ6IHNlcnZpY2UKICByZWNvcmRfdHlwZXM6IERBRU1PTl9ST1RBVEUKLQogIGFjdGlvbjogc3RhcnRlZC1hdWRpdAogIG9iamVjdDoKICAgIHdoYXQ6IHNlcnZpY2UKICByZWNvcmRfdHlwZXM6IERBRU1PTl9TVEFSVAotCiAgYWN0aW9uOiBkZWxldGVkLWdyb3VwLWFjY291bnQtZnJvbQogIG9iamVjdDoKICAgIHByaW1hcnk6IFtpZCwgYWNjdF0KICAgIHdoYXQ6IGFjY291bnQKICByZWNvcmRfdHlwZXM6IERFTF9HUk9VUAotCiAgYWN0aW9uOiBkZWxldGVkLXVzZXItYWNjb3VudAogIG9iamVjdDoKICAgIHByaW1hcnk6IFtpZCwgYWNjdF0KICAgIHdoYXQ6IGFjY291bnQKICByZWNvcmRfdHlwZXM6IERFTF9VU0VSCi0KICBhY3Rpb246IGNoYW5nZWQtYXVkaXQtZmVhdHVyZQogIG9iamVjdDoKICAgIHByaW1hcnk6IGZlYXR1cmUKICAgIHdoYXQ6IHN5c3RlbQogIHJlY29yZF90eXBlczogRkVBVFVSRV9DSEFOR0UKLQogIGFjdGlvbjogcmVsYWJlbGVkLWZpbGVzeXN0ZW0KICByZWNvcmRfdHlwZXM6IEZTX1JFTEFCRUwKLQogIGFjdGlvbjogYXV0aGVudGljYXRlZC10by1ncm91cAogIHJlY29yZF90eXBlczogR1JQX0FVVEgKLQogIDw8OiAqbWFjcm8tdXNlci1zZXNzaW9uCiAgYWN0aW9uOiBjaGFuZ2VkLWdyb3VwLXBhc3N3b3JkCiAgb2JqZWN0OgogICAgcHJpbWFyeTogYWNjdAogICAgd2hhdDogdXNlci1zZXNzaW9uCiAgcmVjb3JkX3R5cGVzOiBHUlBfQ0hBVVRIVE9LCi0KICBhY3Rpb246IG1vZGlmaWVkLWdyb3VwLWFjY291bnQKICBvYmplY3Q6CiAgICBwcmltYXJ5OiBbaWQsIGFjY3RdCiAgICB3aGF0OiBhY2NvdW50CiAgcmVjb3JkX3R5cGVzOiBHUlBfTUdNVAotCiAgYWN0aW9uOiBpbml0aWFsaXplZC1hdWRpdC1zdWJzeXN0ZW0KICByZWNvcmRfdHlwZXM6IEtFUk5FTAotCiAgYWN0aW9uOiBtb2RpZmllZC1sZXZlbC1vZgogIG9iamVjdDoKICAgIHByaW1hcnk6IHByaW50ZXIKICAgIHdoYXQ6IHByaW50ZXIKICByZWNvcmRfdHlwZXM6IExBQkVMX0xFVkVMX0NIQU5HRQotCiAgYWN0aW9uOiBvdmVycm9kZS1sYWJlbC1vZgogIG9iamVjdDoKICAgIHdoYXQ6IG1hYy1jb25maWcKICByZWNvcmRfdHlwZXM6IExBQkVMX09WRVJSSURFCi0KICBvYmplY3Q6CiAgICB3aGF0OiBtYWMtY29uZmlnCiAgcmVjb3JkX3R5cGVzOgogIC0gQVVESVRfREVWX0FMTE9DCiAgLSBBVURJVF9ERVZfREVBTExPQwogIC0gQVVESVRfRlNfUkVMQUJFTAogIC0gQVVESVRfVVNFUl9NQUNfUE9MSUNZX0xPQUQKICAtIEFVRElUX1VTRVJfTUFDX0NPTkZJR19DSEFOR0UKLQogIGFjdGlvbjogY2hhbmdlZC1sb2dpbi1pZC10bwogIHN1YmplY3Q6CiAgICBwcmltYXJ5OiBbb2xkX2F1aWQsIG9sZC1hdWlkXQogICAgc2Vjb25kYXJ5OiB1aWQKICBvYmplY3Q6CiAgICBwcmltYXJ5OiBhdWlkCiAgICB3aGF0OiB1c2VyLXNlc3Npb24KICByZWNvcmRfdHlwZXM6IExPR0lOCi0KICBhY3Rpb246IG1hYy1wZXJtaXNzaW9uCiAgcmVjb3JkX3R5cGVzOiBNQUNfQ0hFQ0sKLQogIGFjdGlvbjogY2hhbmdlZC1zZWxpbnV4LWJvb2xlYW4KICBvYmplY3Q6CiAgICBwcmltYXJ5OiBib29sCiAgICB3aGF0OiBtYWMtY29uZmlnCiAgcmVjb3JkX3R5cGVzOiBNQUNfQ09ORklHX0NIQU5HRQotCiAgYWN0aW9uOiBsb2FkZWQtc2VsaW51eC1wb2xpY3kKICBvYmplY3Q6CiAgICB3aGF0OiBtYWMtY29uZmlnCiAgcmVjb3JkX3R5cGVzOiBNQUNfUE9MSUNZX0xPQUQKLQogIGFjdGlvbjogY2hhbmdlZC1zZWxpbnV4LWVuZm9yY2VtZW50CiAgb2JqZWN0OgogICAgcHJpbWFyeTogZW5mb3JjaW5nCiAgICB3aGF0OiBtYWMtY29uZmlnCiAgcmVjb3JkX3R5cGVzOiBNQUNfU1RBVFVTCi0KICBhY3Rpb246IGFzc2lnbmVkLXVzZXItcm9sZS10bwogIG9iamVjdDoKICAgIHByaW1hcnk6IFtpZCwgYWNjdF0KICAgIHdoYXQ6IGFjY291bnQKICByZWNvcmRfdHlwZXM6IFJPTEVfQVNTSUdOCi0KICBhY3Rpb246IG1vZGlmaWVkLXJvbGUKICByZWNvcmRfdHlwZXM6IFJPTEVfTU9ESUZZCi0KICBhY3Rpb246IHJlbW92ZWQtdXNlLXJvbGUtZnJvbQogIG9iamVjdDoKICAgIHByaW1hcnk6IFtpZCwgYWNjdF0KICAgIHdoYXQ6IGFjY291bnQKICByZWNvcmRfdHlwZXM6IFJPTEVfUkVNT1ZFCi0KICBhY3Rpb246IHZpb2xhdGVkLXNlY2NvbXAtcG9saWN5CiAgb2JqZWN0OgogICAgcHJpbWFyeTogc3lzY2FsbAogICAgd2hhdDogcHJvY2VzcwogIHJlY29yZF90eXBlczogU0VDQ09NUAotCiAgYWN0aW9uOiBzdGFydGVkLXNlcnZpY2UKICBvYmplY3Q6CiAgICBwcmltYXJ5OiB1bml0CiAgICB3aGF0OiBzZXJ2aWNlCiAgcmVjb3JkX3R5cGVzOiBTRVJWSUNFX1NUQVJUCi0KICBhY3Rpb246IHN0b3BwZWQtc2VydmljZQogIG9iamVjdDoKICAgIHByaW1hcnk6IHVuaXQKICAgIHdoYXQ6IHNlcnZpY2UKICByZWNvcmRfdHlwZXM6IFNFUlZJQ0VfU1RPUAotCiAgYWN0aW9uOiBib290ZWQtc3lzdGVtCiAgb2JqZWN0OgogICAgd2hhdDogc3lzdGVtCiAgcmVjb3JkX3R5cGVzOiBTWVNURU1fQk9PVAotCiAgYWN0aW9uOiBjaGFuZ2VkLXRvLXJ1bmxldmVsCiAgb2JqZWN0OgogICAgcHJpbWFyeTogbmV3LWxldmVsCiAgICB3aGF0OiBzeXN0ZW0KICByZWNvcmRfdHlwZXM6IFNZU1RFTV9SVU5MRVZFTAotCiAgYWN0aW9uOiBzaHV0ZG93bi1zeXN0ZW0KICBvYmplY3Q6CiAgICB3aGF0OiBzeXN0ZW0KICByZWNvcmRfdHlwZXM6IFNZU1RFTV9TSFVURE9XTgotCiAgYWN0aW9uOiBzZW50LXRlc3QKICByZWNvcmRfdHlwZXM6IFRFU1QKLQogIGFjdGlvbjogdW5rbm93bgogIHJlY29yZF90eXBlczogVFJVU1RFRF9BUFAKLQogIGFjdGlvbjogc2VudC1tZXNzYWdlCiAgb2JqZWN0OgogICAgcHJpbWFyeTogYWRkcgogIHJlY29yZF90eXBlczogVVNFUgotCiAgPDw6ICptYWNyby11c2VyLXNlc3Npb24KICBhY3Rpb246IHdhcy1hdXRob3JpemVkCiAgcmVjb3JkX3R5cGVzOiBVU0VSX0FDQ1QKLQogIDw8OiAqbWFjcm8tdXNlci1zZXNzaW9uCiAgYWN0aW9uOiBhdXRoZW50aWNhdGVkCiAgcmVjb3JkX3R5cGVzOiBVU0VSX0FVVEgKLQogIGFjdGlvbjogYWNjZXNzLXBlcm1pc3Npb24KICByZWNvcmRfdHlwZXM6IFVTRVJfQVZDCi0KICA8PDogKm1hY3JvLXVzZXItc2Vzc2lvbgogIGFjdGlvbjogY2hhbmdlZC1wYXNzd29yZAogIHJlY29yZF90eXBlczogVVNFUl9DSEFVVEhUT0sKLQogIGFjdGlvbjogcmFuLWNvbW1hbmQKICBvYmplY3Q6CiAgICBwcmltYXJ5OiBjbWQKICAgIHdoYXQ6IHByb2Nlc3MKICByZWNvcmRfdHlwZXM6IFVTRVJfQ01ECiAgZGVzY3JpcHRpb246ID4KICAgIFRoZXNlIG1lc3NhZ2VzIGFyZSBmcm9tIHVzZXItc3BhY2UgYXBwcywgbGlrZSBzdWRvLCB0aGF0IGxvZyBjb21tYW5kcwogICAgYmVpbmcgcnVuIGJ5IGEgdXNlci4gVGhlIHVpZCBjb250YWluZWQgaW4gdGhlc2UgbWVzc2FnZXMgaXMgdXNlcidzIFVJRCBhdAogICAgdGhlIHRpbWUgdGhlIGNvbW1hbmQgd2FzIHJ1bi4gSXQgaXMgbm90IHRoZSAidGFyZ2V0IiBVSUQgdXNlZCB0byBydW4gdGhlCiAgICBjb21tYW5kLCB3aGljaCBpcyBub3JtYWxseSByb290LgotCiAgPDw6ICptYWNyby11c2VyLXNlc3Npb24KICBhY3Rpb246IGVuZGVkLXNlc3Npb24KICByZWNvcmRfdHlwZXM6IFVTRVJfRU5ECi0KICA8PDogKm1hY3JvLXVzZXItc2Vzc2lvbgogIGFjdGlvbjogZXJyb3IKICByZWNvcmRfdHlwZXM6IFVTRVJfRVJSCi0KICA8PDogKm1hY3JvLXVzZXItc2Vzc2lvbgogIGFjdGlvbjogbG9nZ2VkLWluCiAgcmVjb3JkX3R5cGVzOiBVU0VSX0xPR0lOCi0KICA8PDogKm1hY3JvLXVzZXItc2Vzc2lvbgogIGFjdGlvbjogbG9nZ2VkLW91dAogIHJlY29yZF90eXBlczogVVNFUl9MT0dPVVQKLQogIGFjdGlvbjogY2hhbmdlZC1tYWMtY29uZmlndXJhdGlvbgogIHJlY29yZF90eXBlczogVVNFUl9NQUNfQ09ORklHX0NIQU5HRQotCiAgYWN0aW9uOiBsb2FkZWQtbWFjLXBvbGljeQogIHJlY29yZF90eXBlczogVVNFUl9NQUNfUE9MSUNZX0xPQUQKLQogIDw8OiAqbWFjcm8tdXNlci1zZXNzaW9uCiAgYWN0aW9uOiBtb2RpZmllZC11c2VyLWFjY291bnQKICByZWNvcmRfdHlwZXM6IFVTRVJfTUdNVAotCiAgPDw6ICptYWNyby11c2VyLXNlc3Npb24KICBhY3Rpb246IGNoYW5nZWQtcm9sZS10bwogIG9iamVjdDoKICAgIHByaW1hcnk6IHNlbGVjdGVkLWNvbnRleHQKICAgIHdoYXQ6IHVzZXItc2Vzc2lvbgogIHJlY29yZF90eXBlczogVVNFUl9ST0xFX0NIQU5HRQotCiAgYWN0aW9uOiBhY2Nlc3MtZXJyb3IKICByZWNvcmRfdHlwZXM6IFVTRVJfU0VMSU5VWF9FUlIKLQogIDw8OiAqbWFjcm8tdXNlci1zZXNzaW9uCiAgYWN0aW9uOiBzdGFydGVkLXNlc3Npb24KICByZWNvcmRfdHlwZXM6IFVTRVJfU1RBUlQKLQogIGFjdGlvbjogY2hhbmdlZC1jb25maWd1cmF0aW9uCiAgb2JqZWN0OgogICAgcHJpbWFyeTogb3AKICAgIHdoYXQ6IHN5c3RlbQogIHJlY29yZF90eXBlczogVVNZU19DT05GSUcKLQogIGFjdGlvbjogaXNzdWVkLXZtLWNvbnRyb2wKICBvYmplY3Q6CiAgICBwcmltYXJ5OiBvcAogICAgc2Vjb25kYXJ5OiB2bQogICAgd2hhdDogdmlydHVhbC1tYWNoaW5lCiAgcmVjb3JkX3R5cGVzOiBWSVJUX0NPTlRST0wKLQogIGFjdGlvbjogY3JlYXRlZC12bS1pbWFnZQogIHJlY29yZF90eXBlczogVklSVF9DUkVBVEUKLQogIGFjdGlvbjogZGVsZXRlZC12bS1pbWFnZQogIHJlY29yZF90eXBlczogVklSVF9ERVNUUk9ZCi0KICBhY3Rpb246IGNoZWNrZWQtaW50ZWdyaXR5LW9mCiAgcmVjb3JkX3R5cGVzOiBWSVJUX0lOVEVHUklUWV9DSEVDSwotCiAgYWN0aW9uOiBhc3NpZ25lZC12bS1pZAogIG9iamVjdDoKICAgIHByaW1hcnk6IHZtCiAgICB3aGF0OiB2aXJ0dWFsLW1hY2hpbmUKICByZWNvcmRfdHlwZXM6IFZJUlRfTUFDSElORV9JRAotCiAgYWN0aW9uOiBtaWdyYXRlZC12bS1mcm9tCiAgcmVjb3JkX3R5cGVzOiBWSVJUX01JR1JBVEVfSU4KLQogIGFjdGlvbjogbWlncmF0ZWQtdm0tdG8KICByZWNvcmRfdHlwZXM6IFZJUlRfTUlHUkFURV9PVVQKLQogIGFjdGlvbjogYXNzaWduZWQtdm0tcmVzb3VyY2UKICBvYmplY3Q6CiAgICBwcmltYXJ5OiByZXNyYwogICAgc2Vjb25kYXJ5OiB2bQogICAgd2hhdDogdmlydHVhbC1tYWNoaW5lCiAgcmVjb3JkX3R5cGVzOiBWSVJUX1JFU09VUkNFCi0gYWN0aW9uOiB0eXBlZAogIG9iamVjdDoKICAgIHByaW1hcnk6IGRhdGEKICAgIHdoYXQ6IGtleXN0cm9rZXMKICBob3c6IFtjb21tLCBleGVdCiAgcmVjb3JkX3R5cGVzOgogIC0gVFRZCiAgLSBVU0VSX1RUWQo=")
		assets["normalizationData"] = value
	}

	if value, found := assets[key]; found {
		return value, nil
	}
	return nil, fmt.Errorf("asset not found for key=%v", key)
}