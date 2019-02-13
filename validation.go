package quatrix

import (
    "fmt"
    "reflect"
    "strings"
)

// Validate IdsReq
func (s IdsReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate SettingsSetReq
func (s SettingsSetReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate SshKeyCreateReq
func (s SshKeyCreateReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate SessionLoginPostResp
func (s SessionLoginPostResp) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate PfcreateReq
func (s PfcreateReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate QuicklinkCreateReq
func (s QuicklinkCreateReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate SshKeyDeleteReq
func (s SshKeyDeleteReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate ShareRequestReq
func (s ShareRequestReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate FileModifyReq
func (s FileModifyReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate CopyMoveFilesReq
func (s CopyMoveFilesReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate BillingUpgradeReq
func (s BillingUpgradeReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate KeyRequestRespondReq
func (s KeyRequestRespondReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate ResetPasswordResetReq
func (s ResetPasswordResetReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate ProfileRemoveMfaReq
func (s ProfileRemoveMfaReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate ProfileSetReq
func (s ProfileSetReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate UserSignupReq
func (s UserSignupReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate ShareSendRequestReq
func (s ShareSendRequestReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate FileMetadataPostReq
func (s FileMetadataPostReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate UserEditReq
func (s UserEditReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate SshKeyEditReq
func (s SshKeyEditReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate FilesReturnMakedirReq
func (s FilesReturnMakedirReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate PfaddUsersReq
func (s PfaddUsersReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate FilesReturnUploadLinkReq
func (s FilesReturnUploadLinkReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate FileRenameReq
func (s FileRenameReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate ContactCreateReq
func (s ContactCreateReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate ContactEditResp
func (s ContactEditResp) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate FileAddTagReq
func (s FileAddTagReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate PfSetUsersReq
func (s PfSetUsersReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate ShareCreateReq
func (s ShareCreateReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate AutomationEditReq
func (s AutomationEditReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate MakeDirReq
func (s MakeDirReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate UserRemoveMfaReq
func (s UserRemoveMfaReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate UploadLinkReq
func (s UploadLinkReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate ProfileSetPasswordReq
func (s ProfileSetPasswordReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate PfeditUsersReq
func (s PfeditUsersReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate WidgetUploadLinkReq
func (s WidgetUploadLinkReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate ShareDownloadLinkReq
func (s ShareDownloadLinkReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate PgpCreateReq
func (s PgpCreateReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate SearchReq
func (s SearchReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate PgpEditReq
func (s PgpEditReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate ProfileSetMfaReq
func (s ProfileSetMfaReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate UserCreateReq
func (s UserCreateReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate UserDeleteReq
func (s UserDeleteReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate ShareLoginPinReq
func (s ShareLoginPinReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate ResetPasswordRequestReq
func (s ResetPasswordRequestReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate PfUsersListReq
func (s PfUsersListReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate PfdeleteUsersReq
func (s PfdeleteUsersReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate AutomationCreateReq
func (s AutomationCreateReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate QuicklinkLoginPinReq
func (s QuicklinkLoginPinReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate UserSetMfaReq
func (s UserSetMfaReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate UnblockCaptchaReq
func (s UnblockCaptchaReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

// Validate FilesReturnSendReq
func (s FilesReturnSendReq) Validate() error {

    if err := validate(&s); err != nil {
    		return err
    }

	return nil
}

func validate(object interface{}) error {
	val := reflect.ValueOf(object).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		tags := strings.Split(tag.Get("json"), ",")
		if len(tags) == 1 {

			if isEmpty(valueField.Interface()) {
				return fmt.Errorf("empty required filed %s", typeField.Name)
			}
		}
	}
	return nil
}

func isEmpty(object interface{}) bool {

	//First check normal definitions of empty
	if object == nil {
		return true
	} else if object == "" {
		return true
	} else if object == false {
		return true
	}

	switch reflect.TypeOf(object).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(object).Len()
		if s == 0 {
			return true
		}
	case reflect.Struct:
		empty := reflect.New(reflect.TypeOf(object)).Elem().Interface()
		if reflect.DeepEqual(object, empty) {
			return true
		}
	}

	return false
}
