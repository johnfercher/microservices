USE UserDb;

SET character_set_client = utf8;
SET character_set_connection = utf8;
SET character_set_results = utf8;
SET collation_connection = utf8_general_ci;

INSERT INTO user_type_domain (`id`, `name`) VALUES (0, 'person');
INSERT INTO user_type_domain (`id`, `name`) VALUES (1, 'deliverer');
INSERT INTO user_type_domain (`id`, `name`) VALUES (2, 'restaurant');
INSERT INTO user_type_domain (`id`, `name`) VALUES (3, 'market');