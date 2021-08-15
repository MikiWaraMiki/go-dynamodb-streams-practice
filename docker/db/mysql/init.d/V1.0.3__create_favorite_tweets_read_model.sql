CREATE TABLE
IF NOT EXISTS `examples`.`favorite_tweets`
(
	id bigint NOT NULL PRIMARY KEY AUTO_INCREMENT,
	user_uuid VARCHAR (36) NOT NULL,
	tweet_id VARCHAR (36) NOT NULL,
	content VARCHAR (256) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT fk_favorite_tweets_required_users_uuid 
  	FOREIGN KEY (user_uuid) 
  	REFERENCES users(uuid) 
  	ON DELETE RESTRICT
) ENGINE = InnoDB;
