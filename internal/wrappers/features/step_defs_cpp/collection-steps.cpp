#include "steps.hpp"

// Anonymous namespace to handle a linker error
// see https://stackoverflow.com/questions/14320148/linker-error-on-cucumber-cpp-when-dealing-with-multiple-feature-files
namespace {
  WHEN("^I create a collection \'([^\"]*)\'$")
  {
    REGEX_PARAM(std::string, collection_id);

    ScenarioScope<KuzzleCtx> ctx;

    try {
      ctx->kuzzle->collection->create(ctx->index, collection_id);
    } catch (KuzzleException e) {
      BOOST_FAIL(e.getMessage());
    }
  }

  THEN("^the collection \'([^\"]*)\' should exists$")
  {
    REGEX_PARAM(std::string, collection_id);

    ScenarioScope<KuzzleCtx> ctx;

    BOOST_CHECK(ctx->kuzzle->collection->exists(ctx->index, collection_id) == true);
  }

  WHEN("^I check if the collection \'([^\"]*)\' exists$")
  {
    REGEX_PARAM(std::string, collection_id);

    ScenarioScope<KuzzleCtx> ctx;

    ctx->success = ctx->kuzzle->collection->exists(ctx->index, collection_id);
  }

  THEN("^the collection should exist$")
  {
    ScenarioScope<KuzzleCtx> ctx;

    BOOST_CHECK(ctx->success == true);

    ctx->success = false;
  }

  GIVEN("^it has a collection \'([^\"]*)\'$")
  {
    REGEX_PARAM(std::string, collection_name);

    ScenarioScope<KuzzleCtx> ctx;
    ctx->collection = collection_name;

    K_LOG_D("Creating collection: %s", collection_name.c_str());
    try {
      ctx->kuzzle->collection->create(ctx->index, ctx->collection);
    } catch (KuzzleException e) {
      K_LOG_E(e.getMessage().c_str());
      BOOST_FAIL(e.getMessage());
    }
  }

  WHEN("^I list the collections of \'([^\"]*)\'$")
  {
    REGEX_PARAM(std::string, index_name);

    ScenarioScope<KuzzleCtx> ctx;

    ctx->content = ctx->kuzzle->collection->list(index_name);

    json_spirit::Value collectionsList;
    json_spirit::read(ctx->content, collectionsList);
    json_spirit::Value field_value = json_spirit::find_value(collectionsList.get_obj(), "collections");

    ctx->hits = field_value.get_array().size();     
  }

  GIVEN("^the collection has a document with id \'([^\"]*)\'$")
  {
    REGEX_PARAM(std::string, document_id);

    ScenarioScope<KuzzleCtx> ctx;

    query_options options = {0};
    options.refresh = (char*)"wait_for";

    ctx->kuzzle->document->create(ctx->index, ctx->collection, document_id, "{\"a\":\"document\"}");
  }

  WHEN("^I truncate the collection \'([^\"]*)\'$")
  {
    REGEX_PARAM(std::string, collection_id);

    ScenarioScope<KuzzleCtx> ctx;

    query_options options = {0};
    options.refresh = (char*)"wait_for";

    ctx->kuzzle->collection->truncate(ctx->index, collection_id, &options);
  }

  THEN("^the collection \'([^\"]*)\' shall be empty$")
  {
    REGEX_PARAM(std::string, collection_id);

    ScenarioScope<KuzzleCtx> ctx;

    search_result *result = ctx->kuzzle->document->search(ctx->index, collection_id, "{}");

    BOOST_CHECK(result->total == 0);

    kuzzle_free_search_result(result);
  }

  WHEN("^I update the mapping of collection \'([^\"]*)\'$")
  {
    REGEX_PARAM(std::string, collection_id);

    ScenarioScope<KuzzleCtx> ctx;

    std::string mapping = "{\"properties\": {\"gordon\": {\"type\": \"keyword\"}}}";

    ctx->kuzzle->collection->updateMapping(ctx->index, collection_id, mapping);
  }

  THEN("^the mapping of \'([^\"]*)\' should be updated$")
  {
    REGEX_PARAM(std::string, collection_id);

    ScenarioScope<KuzzleCtx> ctx;

    std::string mapping = ctx->kuzzle->collection->getMapping(ctx->index, collection_id);
    std::string expected_mapping = "{\"" + ctx->index + "\":{\"mappings\":{\"" + collection_id + "\":{\"properties\":{\"a\":{\"type\":\"text\",\"fields\":{\"keyword\":{\"type\":\"keyword\",\"ignore_above\":256}}},\"foo\":{\"type\":\"text\",\"fields\":{\"keyword\":{\"type\":\"keyword\",\"ignore_above\":256}}},\"gordon\":{\"type\":\"keyword\"}}}}}}";

    BOOST_CHECK(mapping == expected_mapping);
  }

  WHEN("^I update the specifications of the collection \'([^\"]*)\'$")
  {
    REGEX_PARAM(std::string, collection_id);

    ScenarioScope<KuzzleCtx> ctx;

    std::string specifications = "{\"" + ctx->index + "\":{\""+ collection_id +"\":{\"strict\":false}}}";

    ctx->kuzzle->collection->updateSpecifications(ctx->index, collection_id, specifications);
  }

  THEN("^the specifications of \'([^\"]*)\' should be updated$")
  {
    REGEX_PARAM(std::string, collection_id);

    ScenarioScope<KuzzleCtx> ctx;

    std::string specifications = ctx->kuzzle->collection->getSpecifications(ctx->index, collection_id);
    std::string expected_specifications = "{\"validation\":{\"strict\":false},\"index\":\"" + ctx->index + "\",\"collection\":\"" + collection_id + "\"}";

    BOOST_CHECK(specifications == expected_specifications);

    ctx->kuzzle->collection->updateSpecifications(ctx->index, collection_id, "{\"" + ctx->index + "\":{\""+ collection_id +"\":{\"strict\":true}}}");
  }

  WHEN("^I validate the specifications of \'([^\"]*)\'$")
  {
    REGEX_PARAM(std::string, collection_id);

    ScenarioScope<KuzzleCtx> ctx;

    std::string specifications = "{\"" + ctx->index + "\":{\"" + collection_id + "\":{\"strict\":true}}}";

    ctx->success = ctx->kuzzle->collection->validateSpecifications(specifications);
  }

  THEN("^they should be validated$")
  {
    ScenarioScope<KuzzleCtx> ctx;

    BOOST_CHECK(ctx->success == true);

    ctx->success = false;
  }

  GIVEN("^has specifications$")
  {
    ScenarioScope<KuzzleCtx> ctx;

    std::string specifications = "{\"" + ctx->index + "\":{\""+ ctx->collection +"\":{\"strict\":true}}}";

    ctx->kuzzle->collection->updateSpecifications(ctx->index, ctx->collection, specifications);
  }

  WHEN("^I delete the specifications of \'([^\"]*)\'$")
  {
    REGEX_PARAM(std::string, collection_id);

    ScenarioScope<KuzzleCtx> ctx;

    ctx->kuzzle->collection->deleteSpecifications(ctx->index, collection_id);
  }

  THEN("^the specifications of \'([^\"]*)\' must not exist$")
  {
    REGEX_PARAM(std::string, collection_id);

    ScenarioScope<KuzzleCtx> ctx;

    BOOST_CHECK(ctx->kuzzle->collection->getSpecifications(ctx->index, collection_id) == "");
  }
}