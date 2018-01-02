%typemap(javaimports) kuzzleio::Kuzzle "
/* The type Kuzzle. */"

%javamethodmodifiers kuzzleio::Kuzzle::Kuzzle(const std::string&) "
  /**
   * Constructor
   *
   * @param host - Target Kuzzle host name or IP address
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::Kuzzle(const std::string&, options*) "
  /**
   * Constructor
   *
   * @param host - Target Kuzzle host name or IP address
   * @param options - Request options
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::checkToken(const std::string& token) "
  /**
   * Check an authentication token validity
   *
   * @param token - Token to check (JWT)
   * @return a TokenValidity object
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::connect() "
  /**
   * Connects to a Kuzzle instance using the provided host and port.
   *
   * @return a string which represent an error or null
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::createIndex(const std::string&, query_options*) "
  /**
   * Create a new data index
   *
   * @param index - index name to create
   * @param options - Request options
   * @return a BoolResult object
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::createIndex(const std::string&) "
  /**
   * {@link #createIndex(String, QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::createMyCredentials(const std::string&, json_object*, query_options*) "
  /**
   * Create credentials of the specified strategy for the current user.
   *
   * @param strategy - impacted strategy name
   * @param credentials - credentials to create
   * @param options - Request options
   * @return a JsonResult object
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::createMyCredentials(const std::string&, json_object*) "
  /**
   * {@link #createIndex(String, QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::deleteMyCredentials(const std::string&, query_options*) "
  /**
   * Delete credentials of the specified strategy for the current user.
   *
   * @param strategy- Name of the strategy to remove
   * @param options - Request options
   * @return a BoolResult object
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::deleteMyCredentials(const std::string&) "
  /**
   * {@link #deleteMyCredentials(String, QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::getMyCredentials(const std::string& strategy, query_options *options) "
  /**
   * Get credential information of the specified strategy for the current user.
   *
   * @param strategy - Strategy name to get
   * @param options - Request options
   * @return a JsonResult
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::getMyCredentials(const std::string& strategy) "
  /**
   * {@link #getMyCredentials(String, QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::updateMyCredentials(const std::string& strategy, json_object* credentials, query_options *options) "
  /**
   * Update credentials of the specified strategy for the current user.
   *
   * @param strategy - Strategy name to update
   * @param credentials - Updated credentials content
   * @param options - Request options
   * @return a JsonResult
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::updateMyCredentials(const std::string& strategy, json_object* credentials) "
  /**
   * {@link #updateMyCredentials(String, JsonObject, QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::validateMyCredentials(const std::string& strategy, json_object* credentials, query_options* options) "
  /**
   * Validate credentials of the specified strategy for the current user.
   *
   * @param strategy - Strategy name to validate
   * @param credentials - Credentials content
   * @param options - Request options
   * @return a Bool result
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::validateMyCredentials(const std::string& strategy, json_object* credentials) "
  /**
   * {@link #validateMyCredentials(String, JsonObject, QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::logout() "
  /**
   * Logout method
   *
   * @param listener - Response callback listener
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::login(const std::string&, json_object*, int) "
  /**
   * Log-in Strategy name to use for the authentication
   *
   * @param strategy - Strategy name to use for the authentication
   * @param credentials - Login credentials
   * @param expiresIn - Token expiration delay
   * @return StringResult
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::login(const std::string&, json_object*) "
  /**
   * Log-in Strategy name to use for the authentication
   *
   * @param strategy - Strategy name to use for the authentication
   * @param credentials - Login credentials
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::login(const std::string&) "
  /**
   * Log-in Strategy name to use for the authentication
   *
   * @param strategy - Strategy name to use for the authentication
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::login(const std::string&) "
  /**
   * Get all Kuzzle usage statistics frames
   *
   * @param options - Request options
   * @param listener - Response callback listener
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::getAllStatistics(query_options*) "
  /**
   * Get all Kuzzle usage statistics frames
   *
   * @param options - Request options
   * @return a AllStatisticsResult
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::getAllStatistics() "
  /**
   * {@link #getAllStatistics(QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::getStatistics(unsigned long, query_options*) "
  /**
   * Get Kuzzle usage statistics
   *
   * @param options - Request options
   * @return a StatisticsResult
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::getStatistics(unsigned long) "
  /**
   * {@link #getStatistics(QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::getAutoRefresh(const std::string&, query_options*) "
  /**
   * Gets the autoRefresh value for the provided data index name
   *
   * @param index - Data index name
   * @param options - Request options
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::getAutoRefresh(const std::string&) "
  /**
   * {@link #getAutoRefresh(String, QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::getAutoRefresh() "
  /**
   * {@link #getAutoRefresh(String, QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::getJwt() "
  /**
   * Authentication token getter
   *
   * @return a string which is the jwt
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::getMyRights(query_options*) "
  /**
   * Gets the rights array for the currently logged user.
   *
   * @param options - Request options
   * @return a JsonResult
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::getMyRights() "
  /**
   * {@link #getMyRights(QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::getServerInfo(query_options*) "
  /**
   * Gets server info.
   *
   * @param options - Request options
   * #return a JsonResult
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::getServerInfo() "
  /**
   * {@link #getServerInfo(QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::listCollections(const std::string&, query_options*) "
  /**
   * List data collections
   *
   * @param index - Parent data index name
   * @param options - Request options
   * @return a CollectionListResult
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::listCollections(const std::string&) "
  /**
   * {@link #listCollections(String, QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::listCollections() "
  /**
   * {@link #listCollections(String, QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::listIndexes(query_options*) "
  /**
   * List data indexes
   *
   * @param options - Request options
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::listIndexes() "
  /**
   * {@link #listIndexes(QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::disconnect() "
  /**
   * Disconnect from Kuzzle and invalidate this instance.
   * Does not fire a disconnected event.
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::logout() "
  /**
   * Logout method
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::removeListener(enum Event) "
  /**
   * Removes a listener from an event.
   *
   * @param event - Event name
   * @return this
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::query(kuzzle_request*, query_options*) "
  /**
   * This is a low-level method, exposed to allow advanced SDK users to bypass high-level methods.
   * Base method used to send queries to Kuzzle
   *
   * @param query - Query content
   * @param options - Request options
   * @return a KuzzleResponse
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::query(kuzzle_request*) "
  /**
   * {@link #query(KuzzleRequest, QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::refreshIndex(const std::string& index, query_options* options) "
  /**
   * Forces the default data index to refresh on each modification
   *
   * @param options - Request options
   * @return this
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::refreshIndex(const std::string& index) "
  /**
   * {@link #refreshIndex(String, QueryOptions)}
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::removeListener(enum Event) "
  /**
   * Removes a listener from an event.
   *
   * @param event - Event name
   * @return this
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::setDefaultIndex(const std::string& index) "
  /**
   * Default index setter
   *
   * @param index - New default index name
   * @return this
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::replayQueue() "
  /**
   * Replays the requests queued during offline mode.
   * Works only if the SDK is not in a disconnected state, and if the autoReplay option is set to false.
   *
   * @return this
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::setAutoReplay(bool autoReplay) "
  /**
   * autoReplay option setter
   *
   * @param autoReplay - New autoReplay option value
   * @return this
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::setJwt(const std::string& jwt) "
  /**
   * Set a new JWT and trigger the 'loginAttempt' event.
   *
   * @param jwt - New authentication JSON Web Token
   * @return this 
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::startQueuing() "
  /**
   * Starts the requests queuing. Works only during offline mode, and if the autoQueue option is set to false.
   *
   * @return this
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::stopQueuing() "
  /**
   * Stops the requests queuing. Works only during offline mode, and if the autoQueue option is set to false.
   *
   * @return this
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::unsetJwt() "
  /**
   * Unset the authentication token and cancel all subscriptions
   *
   * @return this
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::updateSelf(user_data* content, query_options* options) "
  /**
   * Update the currently authenticated user informations
   *
   * @param content - Current user infos to update
   * @param options - Request options
   * @return updated user
   */
  public";

%javamethodmodifiers kuzzleio::Kuzzle::updateSelf(user_data* content) "
  /**
   * {@link #updateSelf(UserData, Options)}
   */
  public";