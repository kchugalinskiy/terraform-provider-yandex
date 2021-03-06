package yandex

import "github.com/hashicorp/terraform/helper/schema"

func resourceYandexIAMServiceAccountIAMMember() *schema.Resource {
	return resourceIamMemberWithImport(IamServiceAccountSchema, newServiceAccountIamUpdater, serviceAccountIDParseFunc)
}
