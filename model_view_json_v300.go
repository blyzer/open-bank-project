/*
 * OpenBankProject-User
 *
 * An Open Source API for Banks. (c) TESOBE Ltd. 2011 - 2018. Licensed under the AGPL and commercial licences.
 *
 * API version: v3_1_0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ViewJsonV300 struct {
	CanSeeTransactionStartDate bool `json:"can_see_transaction_start_date,omitempty"`
	CanAddUrl bool `json:"can_add_url,omitempty"`
	CanAddWhereTag bool `json:"can_add_where_tag,omitempty"`
	CanSeeTransactionThisBankAccount bool `json:"can_see_transaction_this_bank_account,omitempty"`
	CanSeeBankAccountOwners bool `json:"can_see_bank_account_owners,omitempty"`
	CanSeeBankRoutingAddress bool `json:"can_see_bank_routing_address,omitempty"`
	CanSeePrivateAlias bool `json:"can_see_private_alias,omitempty"`
	CanEditOwnerComment bool `json:"can_edit_owner_comment,omitempty"`
	CanSeeOtherAccountNationalIdentifier bool `json:"can_see_other_account_national_identifier,omitempty"`
	CanSeeBankRoutingScheme bool `json:"can_see_bank_routing_scheme,omitempty"`
	CanSeePublicAlias bool `json:"can_see_public_alias,omitempty"`
	CanSeePhysicalLocation bool `json:"can_see_physical_location,omitempty"`
	CanSeeOwnerComment bool `json:"can_see_owner_comment,omitempty"`
	CanSeeBankAccountIban bool `json:"can_see_bank_account_iban,omitempty"`
	CanSeeCorporateLocation bool `json:"can_see_corporate_location,omitempty"`
	CanSeeBankAccountNumber bool `json:"can_see_bank_account_number,omitempty"`
	CanSeeOtherAccountBankName bool `json:"can_see_other_account_bank_name,omitempty"`
	Description string `json:"description,omitempty"`
	CanSeeBankAccountRoutingScheme bool `json:"can_see_bank_account_routing_scheme,omitempty"`
	CanSeeTransactionOtherBankAccount bool `json:"can_see_transaction_other_bank_account,omitempty"`
	CanDeleteCorporateLocation bool `json:"can_delete_corporate_location,omitempty"`
	CanSeeComments bool `json:"can_see_comments,omitempty"`
	CanSeeBankAccountBankName bool `json:"can_see_bank_account_bank_name,omitempty"`
	CanAddMoreInfo bool `json:"can_add_more_info,omitempty"`
	CanSeeOtherAccountNumber bool `json:"can_see_other_account_number,omitempty"`
	CanSeeOtherAccountSwiftBic bool `json:"can_see_other_account_swift_bic,omitempty"`
	CanAddOpenCorporatesUrl bool `json:"can_add_open_corporates_url,omitempty"`
	CanSeeOtherAccountKind bool `json:"can_see_other_account_kind,omitempty"`
	CanAddTransactionRequestToOwnAccount bool `json:"can_add_transaction_request_to_own_account,omitempty"`
	CanDeletePhysicalLocation bool `json:"can_delete_physical_location,omitempty"`
	CanSeeBankAccountLabel bool `json:"can_see_bank_account_label,omitempty"`
	CanSeeTransactionCurrency bool `json:"can_see_transaction_currency,omitempty"`
	IsPublic bool `json:"is_public,omitempty"`
	CanSeeTransactionFinishDate bool `json:"can_see_transaction_finish_date,omitempty"`
	CanSeeBankAccountRoutingAddress bool `json:"can_see_bank_account_routing_address,omitempty"`
	CanAddTag bool `json:"can_add_tag,omitempty"`
	CanSeeImages bool `json:"can_see_images,omitempty"`
	CanSeeBankAccountCreditLimit bool `json:"can_see_bank_account_credit_limit,omitempty"`
	CanSeeBankAccountCurrency bool `json:"can_see_bank_account_currency,omitempty"`
	HideMetadataIfAliasUsed bool `json:"hide_metadata_if_alias_used,omitempty"`
	CanDeleteWhereTag bool `json:"can_delete_where_tag,omitempty"`
	Alias string `json:"alias,omitempty"`
	CanAddImageUrl bool `json:"can_add_image_url,omitempty"`
	CanAddComment bool `json:"can_add_comment,omitempty"`
	CanSeeImageUrl bool `json:"can_see_image_url,omitempty"`
	Id string `json:"id,omitempty"`
	CanSeeBankAccountNationalIdentifier bool `json:"can_see_bank_account_national_identifier,omitempty"`
	CanAddCounterparty bool `json:"can_add_counterparty,omitempty"`
	CanAddTransactionRequestToAnyAccount bool `json:"can_add_transaction_request_to_any_account,omitempty"`
	CanSeeTags bool `json:"can_see_tags,omitempty"`
	CanSeeOpenCorporatesUrl bool `json:"can_see_open_corporates_url,omitempty"`
	ShortName string `json:"short_name,omitempty"`
	CanDeleteTag bool `json:"can_delete_tag,omitempty"`
	CanSeeOtherAccountRoutingScheme bool `json:"can_see_other_account_routing_scheme,omitempty"`
	CanSeeMoreInfo bool `json:"can_see_more_info,omitempty"`
	CanSeeTransactionMetadata bool `json:"can_see_transaction_metadata,omitempty"`
	CanDeleteComment bool `json:"can_delete_comment,omitempty"`
	CanSeeWhereTag bool `json:"can_see_where_tag,omitempty"`
	CanAddPrivateAlias bool `json:"can_add_private_alias,omitempty"`
	CanAddPublicAlias bool `json:"can_add_public_alias,omitempty"`
	CanSeeBankAccountSwiftBic bool `json:"can_see_bank_account_swift_bic,omitempty"`
	CanAddImage bool `json:"can_add_image,omitempty"`
	CanSeeTransactionType bool `json:"can_see_transaction_type,omitempty"`
	CanSeeOtherAccountRoutingAddress bool `json:"can_see_other_account_routing_address,omitempty"`
	CanSeeOtherAccountIban bool `json:"can_see_other_account_iban,omitempty"`
	CanAddPhysicalLocation bool `json:"can_add_physical_location,omitempty"`
	CanAddCorporateLocation bool `json:"can_add_corporate_location,omitempty"`
	CanDeleteImage bool `json:"can_delete_image,omitempty"`
	CanSeeUrl bool `json:"can_see_url,omitempty"`
	CanSeeBankAccountBalance bool `json:"can_see_bank_account_balance,omitempty"`
	CanSeeOtherBankRoutingAddress bool `json:"can_see_other_bank_routing_address,omitempty"`
	CanSeeTransactionBalance bool `json:"can_see_transaction_balance,omitempty"`
	MetadataView string `json:"metadata_view,omitempty"`
	CanSeeTransactionAmount bool `json:"can_see_transaction_amount,omitempty"`
	CanSeeOtherAccountMetadata bool `json:"can_see_other_account_metadata,omitempty"`
	CanSeeBankAccountType bool `json:"can_see_bank_account_type,omitempty"`
	CanSeeOtherBankRoutingScheme bool `json:"can_see_other_bank_routing_scheme,omitempty"`
	CanSeeTransactionDescription bool `json:"can_see_transaction_description,omitempty"`
}
