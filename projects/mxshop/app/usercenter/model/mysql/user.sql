CREATE TABLE user (
                      id bigint AUTO_INCREMENT,
                      uid bigint NOT NULL DEFAULT '0' COMMENT 'The user id',
                      name varchar(255) NOT NULL COMMENT 'The username',
                      password varchar(255) NOT NULL DEFAULT '' COMMENT 'The user password',
                      mobile varchar(255) NOT NULL DEFAULT '' COMMENT 'The mobile phone number',
                      gender integer NOT NULL DEFAULT '0' COMMENT 'gender,1 male|2 female|0 unknown',
                      create_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                      update_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                      UNIQUE mobile_index (mobile),
                      UNIQUE name_index (name),
                      PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'user table';