package main

/*
  #cgo CFLAGS: -std=c99 -I../../../headers
  #cgo LDFLAGS: -ljson-c

  #include <stdlib.h>
  #include "kuzzlesdk.h"

  static void free_char_array(char **arr, size_t length) {
    if (arr != NULL) {
      for(int i = 0; i < length; i++) {
        free(arr[i]);
      }

      free(arr);
    }
  }
*/
import "C"

import (
	"unsafe"
)

//export kuzzle_wrapper_free_kuzzle_request
func kuzzle_wrapper_free_kuzzle_request(st *C.kuzzle_request) {
	if st != nil {
		C.free(unsafe.Pointer(st.request_id))
		C.free(unsafe.Pointer(st.controller))
		C.free(unsafe.Pointer(st.action))
		C.free(unsafe.Pointer(st.index))
		C.free(unsafe.Pointer(st.collection))
		C.free(unsafe.Pointer(st.id))
		C.free(unsafe.Pointer(st.scroll))
		C.free(unsafe.Pointer(st.scroll_id))
		C.free(unsafe.Pointer(st.strategy))
		C.free(unsafe.Pointer(st.scope))
		C.free(unsafe.Pointer(st.state))
		C.free(unsafe.Pointer(st.user))
		C.free(unsafe.Pointer(st.member))
		C.free(unsafe.Pointer(st.member1))
		C.free(unsafe.Pointer(st.member2))
		C.free(unsafe.Pointer(st.unit))
		C.free(unsafe.Pointer(st.field))
		C.free(unsafe.Pointer(st.subcommand))
		C.free(unsafe.Pointer(st.pattern))
		C.free(unsafe.Pointer(st.min))
		C.free(unsafe.Pointer(st.max))
		C.free(unsafe.Pointer(st.limit))
		C.free(unsafe.Pointer(st.match))

		C.free_char_array(st.members, st.members_length)
		C.free_char_array(st.keys, st.keys_length)
		C.free_char_array(st.fields, st.fields_length)

		kuzzle_wrapper_free_json_object(st.body)
		kuzzle_wrapper_free_json_object(st.volatiles)
		kuzzle_wrapper_free_json_object(st.options)

		C.free(unsafe.Pointer(st))
	}

}

//export kuzzle_wrapper_free_query_object
func kuzzle_wrapper_free_query_object(st *C.query_object) {
	if st != nil {
		kuzzle_wrapper_free_json_object(st.query)
		C.free(unsafe.Pointer(st.request_id))

		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_offline_queue
func kuzzle_wrapper_free_offline_queue(st *C.offline_queue) {
	if st != nil && st.queries != nil {
		queries := (*[1<<30 - 1]*C.query_object)(unsafe.Pointer(st.queries))[:int(st.queries_length):int(st.queries_length)]

		for _, query := range queries {
			kuzzle_wrapper_free_query_object(query)
		}

		C.free(unsafe.Pointer(st.queries))
	}

	C.free(unsafe.Pointer(st))
}

//export kuzzle_wrapper_free_query_options
func kuzzle_wrapper_free_query_options(st *C.query_options) {
	if st != nil {
		C.free(unsafe.Pointer(st.scroll))
		C.free(unsafe.Pointer(st.scroll_id))
		C.free(unsafe.Pointer(st.refresh))
		C.free(unsafe.Pointer(st.if_exist))
		kuzzle_wrapper_free_json_object(st.volatiles)

		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_room_options
func kuzzle_wrapper_free_room_options(st *C.room_options) {
	if st != nil {
		C.free(unsafe.Pointer(st.scope))
		C.free(unsafe.Pointer(st.state))
		C.free(unsafe.Pointer(st.user))
		kuzzle_wrapper_free_json_object(st.volatiles)
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_options
func kuzzle_wrapper_free_options(st *C.options) {
	if st != nil {
		C.free(unsafe.Pointer(st.refresh))
		C.free(unsafe.Pointer(st.default_index))
		kuzzle_wrapper_free_json_object(st.headers)
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_meta
func kuzzle_wrapper_free_meta(st *C.meta) {
	if st != nil {
		C.free(unsafe.Pointer(st.author))
		C.free(unsafe.Pointer(st.updater))
		C.free(unsafe.Pointer(st))
	}
}

// do not export => used to free the content of a structure
// and not the structure itself
func _free_policy_restriction(st *C.policy_restriction) {
	if st != nil {
		C.free(unsafe.Pointer(st.index))
		C.free_char_array(st.collections, st.collections_length)
	}
}

//export kuzzle_wrapper_free_policy_restriction
func kuzzle_wrapper_free_policy_restriction(st *C.policy_restriction) {
	_free_policy_restriction(st)
	C.free(unsafe.Pointer(st))
}

// do not export => used to free the content of a structure
// and not the structure itself
func _free_policy(st *C.policy) {
	if st != nil {
		C.free(unsafe.Pointer(st.role_id))

		if st.restricted_to != nil {
			restrictions := (*[1<<30 - 1]C.policy_restriction)(unsafe.Pointer(st.restricted_to))[:int(st.restricted_to_length):int(st.restricted_to_length)]

			for _, restriction := range restrictions {
				_free_policy_restriction(&restriction)
			}

			C.free(unsafe.Pointer(st.restricted_to))
		}
	}
}

//export kuzzle_wrapper_free_policy
func kuzzle_wrapper_free_policy(st *C.policy) {
	_free_policy(st)
	C.free(unsafe.Pointer(st))
}

// do not export => used to free the content of a structure
// and not the structure itself
func _free_profile(st *C.profile) {
	if st != nil {
		C.free(unsafe.Pointer(st.id))

		if st.policies != nil {
			policies := (*[1<<30 - 1]C.policy)(unsafe.Pointer(st.policies))[:int(st.policies_length):int(st.policies_length)]

			for _, policy := range policies {
				_free_policy(&policy)
			}

			C.free(unsafe.Pointer(st.policies))
		}
	}
}

//export kuzzle_wrapper_free_profile
func kuzzle_wrapper_free_profile(st *C.profile) {
	_free_profile(st)
	C.free(unsafe.Pointer(st))
}

//do not export
func _free_role(st *C.role) {
	if st != nil {
		C.free(unsafe.Pointer(st.id))
		kuzzle_wrapper_free_json_object(st.controllers)
	}
}

//export kuzzle_wrapper_free_role
func kuzzle_wrapper_free_role(st *C.role) {
	_free_role(st)
	C.free(unsafe.Pointer(st))
}

//do not export
func _free_user(st *C.user) {
	if st != nil {
		C.free(unsafe.Pointer(st.id))
		kuzzle_wrapper_free_json_object(st.content)
		C.free_char_array(st.profile_ids, st.profile_ids_length)
	}
}

//export kuzzle_wrapper_free_user
func kuzzle_wrapper_free_user(st *C.user) {
	_free_user(st)
	C.free(unsafe.Pointer(st))
}

//export kuzzle_wrapper_free_user_data
func kuzzle_wrapper_free_user_data(st *C.user_data) {
	if st != nil {
		kuzzle_wrapper_free_json_object(st.content)
		C.free_char_array(st.profile_ids, st.profile_ids_length)
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_collection
func kuzzle_wrapper_free_collection(st *C.collection) {
	if st != nil {
		C.free(unsafe.Pointer(st.index))
		C.free(unsafe.Pointer(st.collection))
		C.free(unsafe.Pointer(st))
	}
}

//do not export
func _free_document(st *C.document) {
	if st != nil {
		C.free(unsafe.Pointer(st.id))
		C.free(unsafe.Pointer(st.index))
		C.free(unsafe.Pointer(st.shards))
		C.free(unsafe.Pointer(st.result))
		C.free(unsafe.Pointer(st.collection))

		kuzzle_wrapper_free_json_object(st.content)

		kuzzle_wrapper_free_meta(st.meta)
		kuzzle_wrapper_free_collection(st._collection)
	}
}

//export kuzzle_wrapper_free_document
func kuzzle_wrapper_free_document(st *C.document) {
	_free_document(st)
	C.free(unsafe.Pointer(st))
}

//export kuzzle_wrapper_free_document_result
func kuzzle_wrapper_free_document_result(st *C.document_result) {
	if st != nil {
		kuzzle_wrapper_free_document(st.result)
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_notification_content
func kuzzle_wrapper_free_notification_content(st *C.notification_content) {
	if st != nil {
		C.free(unsafe.Pointer(st.id))
		kuzzle_wrapper_free_meta(st.meta)
		kuzzle_wrapper_free_json_object(st.content)
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_notification_result
func kuzzle_wrapper_free_notification_result(st *C.notification_result) {
	if st != nil {
		C.free(unsafe.Pointer(st.request_id))
		C.free(unsafe.Pointer(st.index))
		C.free(unsafe.Pointer(st.collection))
		C.free(unsafe.Pointer(st.controller))
		C.free(unsafe.Pointer(st.action))
		C.free(unsafe.Pointer(st.protocol))
		C.free(unsafe.Pointer(st.scope))
		C.free(unsafe.Pointer(st.state))
		C.free(unsafe.Pointer(st.user))
		C.free(unsafe.Pointer(st.n_type))
		C.free(unsafe.Pointer(st.room_id))
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))

		kuzzle_wrapper_free_json_object(st.volatiles)

		kuzzle_wrapper_free_notification_content(st.result)

		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_profile_result
func kuzzle_wrapper_free_profile_result(st *C.profile_result) {
	if st != nil {
		kuzzle_wrapper_free_profile(st.profile)
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_profiles_result
func kuzzle_wrapper_free_profiles_result(st *C.profiles_result) {
	if st != nil {
		if st.profiles != nil {
			profiles := (*[1<<30 - 1]C.profile)(unsafe.Pointer(st.profiles))[:int(st.profiles_length):int(st.profiles_length)]

			for _, profile := range profiles {
				_free_profile(&profile)
			}

			C.free(unsafe.Pointer(st.profiles))
		}

		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_role_result
func kuzzle_wrapper_free_role_result(st *C.role_result) {
	if st != nil {
		kuzzle_wrapper_free_role(st.role)
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

// do not export => used to free the content of a structure
// and not the structure itself
func _free_user_right(st *C.user_right) {
	if st != nil {
		C.free(unsafe.Pointer(st.controller))
		C.free(unsafe.Pointer(st.action))
		C.free(unsafe.Pointer(st.index))
		C.free(unsafe.Pointer(st.collection))
		C.free(unsafe.Pointer(st.value))
	}
}

//export kuzzle_wrapper_free_user_right
func kuzzle_wrapper_free_user_right(st *C.user_right) {
	_free_user_right(st)
	C.free(unsafe.Pointer(st))
}

//export kuzzle_wrapper_free_user_rights_result
func kuzzle_wrapper_free_user_rights_result(st *C.user_rights_result) {
	if st != nil {
		if st.user_rights != nil {
			rights := (*[1<<30 - 1]C.user_right)(unsafe.Pointer(st.user_rights))[:int(st.user_rights_length):int(st.user_rights_length)]

			for _, right := range rights {
				_free_user_right(&right)
			}

			C.free(unsafe.Pointer(st.user_rights))
		}

		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_user_result
func kuzzle_wrapper_free_user_result(st *C.user_result) {
	if st != nil {
		kuzzle_wrapper_free_user(st.user)
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

// do not export => used to free the content of a structure
// and not the structure itself
func _free_statistics(st *C.statistics) {
	if st != nil {
		kuzzle_wrapper_free_json_object(st.completed_requests)
		kuzzle_wrapper_free_json_object(st.connections)
		kuzzle_wrapper_free_json_object(st.failed_requests)
		kuzzle_wrapper_free_json_object(st.ongoing_requests)
	}
}

//export kuzzle_wrapper_free_statistics
func kuzzle_wrapper_free_statistics(st *C.statistics) {
	_free_statistics(st)
	C.free(unsafe.Pointer(st))
}

//export kuzzle_wrapper_free_statistics_result
func kuzzle_wrapper_free_statistics_result(st *C.statistics_result) {
	if st != nil {
		kuzzle_wrapper_free_statistics(st.result)
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_all_statistics_result
func kuzzle_wrapper_free_all_statistics_result(st *C.all_statistics_result) {
	if st != nil {
		if st.result != nil {
			stats := (*[1<<30 - 1]C.statistics)(unsafe.Pointer(st.result))

			for _, stat := range stats {
				_free_statistics(&stat)
			}

			C.free(unsafe.Pointer(st.result))
		}

		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_geopos_result
func kuzzle_wrapper_free_geopos_result(st *C.geopos_result) {
	if st != nil {
		C.free(unsafe.Pointer(st.result))
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_token_validity
func kuzzle_wrapper_free_token_validity(st *C.token_validity) {
	if st != nil {
		C.free(unsafe.Pointer(st.state))
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_kuzzle_response
func kuzzle_wrapper_free_kuzzle_response(st *C.kuzzle_response) {
	if st != nil {
		C.free(unsafe.Pointer(st.request_id))
		C.free(unsafe.Pointer(st.index))
		C.free(unsafe.Pointer(st.collection))
		C.free(unsafe.Pointer(st.controller))
		C.free(unsafe.Pointer(st.action))
		C.free(unsafe.Pointer(st.room_id))
		C.free(unsafe.Pointer(st.channel))
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))

		kuzzle_wrapper_free_json_object(st.result)
		kuzzle_wrapper_free_json_object(st.volatiles)

		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_json_result
func kuzzle_wrapper_free_json_result(st *C.json_result) {
	if st != nil {
		kuzzle_wrapper_free_json_object(st.result)
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_json_array_result
func kuzzle_wrapper_free_json_array_result(st *C.json_array_result) {
	if st != nil {
		if st.result != nil {
			jobjects := (*[1<<30 - 1]*C.json_object)(unsafe.Pointer(st.result))[:int(st.result_length):int(st.result_length)]

			for _, jobject := range jobjects {
				kuzzle_wrapper_free_json_object(jobject)
			}

			C.free(unsafe.Pointer(st.result))
		}

		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_bool_result
func kuzzle_wrapper_free_bool_result(st *C.bool_result) {
	if st != nil {
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_int_result
func kuzzle_wrapper_free_int_result(st *C.int_result) {
	if st != nil {
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_double_result
func kuzzle_wrapper_free_double_result(st *C.double_result) {
	if st != nil {
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_int_array_result
func kuzzle_wrapper_free_int_array_result(st *C.int_array_result) {
	if st != nil {
		C.free(unsafe.Pointer(st.result))
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_string_result
func kuzzle_wrapper_free_string_result(st *C.string_result) {
	if st != nil {
		C.free(unsafe.Pointer(st.result))
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_string_array_result
func kuzzle_wrapper_free_string_array_result(st *C.string_array_result) {
	if st != nil {
		C.free_char_array(st.result, st.result_length)
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_search_filters
func kuzzle_wrapper_free_search_filters(st *C.search_filters) {
	if st != nil {
		kuzzle_wrapper_free_json_object(st.query)
		kuzzle_wrapper_free_json_object(st.sort)
		kuzzle_wrapper_free_json_object(st.aggregations)
		kuzzle_wrapper_free_json_object(st.search_after)
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_document_search
func kuzzle_wrapper_free_document_search(st *C.document_search) {
	if st != nil {
		C.free(unsafe.Pointer(st.scroll_id))

		if st.hits != nil {
			hits := (*[1<<30 - 1]C.document)(unsafe.Pointer(st.hits))[:int(st.hits_length):int(st.hits_length)]

			for _, document := range hits {
				_free_document(&document)
			}

			C.free(unsafe.Pointer(st.hits))
		}

		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_profile_search
func kuzzle_wrapper_free_profile_search(st *C.profile_search) {
	if st != nil {
		C.free(unsafe.Pointer(st.scroll_id))

		if st.hits != nil {
			hits := (*[1<<30 - 1]C.profile)(unsafe.Pointer(st.hits))[:int(st.hits_length):int(st.hits_length)]

			for _, profile := range hits {
				_free_profile(&profile)
			}

			C.free(unsafe.Pointer(st.hits))
		}

		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_role_search
func kuzzle_wrapper_free_role_search(st *C.role_search) {
	if st != nil {
		if st.hits != nil {
			hits := (*[1<<30 - 1]C.role)(unsafe.Pointer(st.hits))[:int(st.hits_length):int(st.hits_length)]

			for _, role := range hits {
				_free_role(&role)
			}

			C.free(unsafe.Pointer(st.hits))
		}

		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_ack_result
func kuzzle_wrapper_free_ack_result(st *C.ack_result) {
	if st != nil {
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_shards_result
func kuzzle_wrapper_free_shards_result(st *C.shards_result) {
	if st != nil {
		C.free(unsafe.Pointer(st.result))
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_specification
func kuzzle_wrapper_free_specification(st *C.specification) {
	if st != nil {
		kuzzle_wrapper_free_json_object(st.fields)
		kuzzle_wrapper_free_json_object(st.validators)
		C.free(unsafe.Pointer(st))
	}
}

//do not export
func _free_specification_entry(st *C.specification_entry) {
	if st != nil {
		kuzzle_wrapper_free_specification(st.validation)
		C.free(unsafe.Pointer(st.index))
		C.free(unsafe.Pointer(st.collection))
	}
}

//export kuzzle_wrapper_free_specification_entry
func kuzzle_wrapper_free_specification_entry(st *C.specification_entry) {
	_free_specification_entry(st)
	C.free(unsafe.Pointer(st))
}

//export kuzzle_wrapper_free_specification_result
func kuzzle_wrapper_free_specification_result(st *C.specification_result) {
	if st != nil {
		kuzzle_wrapper_free_specification(st.result)
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_search_result
func kuzzle_wrapper_free_search_result(st *C.search_result) {
	if st != nil {
		kuzzle_wrapper_free_document_search(st.result)
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_search_profiles_result
func kuzzle_wrapper_free_search_profiles_result(st *C.search_profiles_result) {
	if st != nil {
		kuzzle_wrapper_free_profile_search(st.result)
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_search_roles_result
func kuzzle_wrapper_free_search_roles_result(st *C.search_roles_result) {
	if st != nil {
		kuzzle_wrapper_free_role_search(st.result)
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_specification_search
func kuzzle_wrapper_free_specification_search(st *C.specification_search) {
	if st != nil {
		if st.hits != nil {
			hits := (*[1<<30 - 1]C.specification_entry)(unsafe.Pointer(st.hits))[:int(st.hits_length):int(st.hits_length)]

			for _, entry := range hits {
				_free_specification_entry(&entry)
			}

			C.free(unsafe.Pointer(st.hits))
			C.free(unsafe.Pointer(st.scroll_id))
			C.free(unsafe.Pointer(st))
		}
	}
}

//export kuzzle_wrapper_free_specification_search_result
func kuzzle_wrapper_free_specification_search_result(st *C.specification_search_result) {
	if st != nil {
		kuzzle_wrapper_free_specification_search(st.result)
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_mapping
func kuzzle_wrapper_free_mapping(st *C.mapping) {
	if st != nil {
		kuzzle_wrapper_free_json_object(st.mapping)
		kuzzle_wrapper_free_collection(st.collection)
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_mapping_result
func kuzzle_wrapper_free_mapping_result(st *C.mapping_result) {
	if st != nil {
		kuzzle_wrapper_free_mapping(st.result)
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_void_result
func kuzzle_wrapper_free_void_result(st *C.void_result) {
	if st != nil {
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//do not export
func _free_collection_entry(st *C.collection_entry) {
	if st != nil {
		C.free(unsafe.Pointer(st.name))
	}
}

//export kuzzle_wrapper_free_collection_entry
func kuzzle_wrapper_free_collection_entry(st *C.collection_entry) {
	_free_collection_entry(st)
	C.free(unsafe.Pointer(st))
}

//export kuzzle_wrapper_free_collection_entry_result
func kuzzle_wrapper_free_collection_entry_result(st *C.collection_entry_result) {
	if st != nil {
		if st.result != nil {
			entries := (*[1<<30 - 1]C.collection_entry)(unsafe.Pointer(st.result))[:int(st.result_length):int(st.result_length)]

			for _, entry := range entries {
				_free_collection_entry(&entry)
			}

			C.free(unsafe.Pointer(st.result))
		}

		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_user_search
func kuzzle_wrapper_free_user_search(st *C.user_search) {
	if st != nil {
		if st.hits != nil {
			hits := (*[1<<30 - 1]C.user)(unsafe.Pointer(st.hits))[:int(st.hits_length):int(st.hits_length)]

			for _, user := range hits {
				_free_user(&user)
			}

			C.free(unsafe.Pointer(st.hits))
		}

		C.free(unsafe.Pointer(st.scroll_id))
		C.free(unsafe.Pointer(st))
	}
}

//export kuzzle_wrapper_free_search_users_result
func kuzzle_wrapper_free_search_users_result(st *C.search_users_result) {
	if st != nil {
		kuzzle_wrapper_free_user_search(st.result)
		C.free(unsafe.Pointer(st.error))
		C.free(unsafe.Pointer(st.stack))
		C.free(unsafe.Pointer(st))
	}
}