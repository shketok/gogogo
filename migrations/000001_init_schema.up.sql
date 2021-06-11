CREATE TABLE IF NOT EXISTS "restaurant"
(
    "id"              bigserial PRIMARY KEY,
    "restaurant_name" varchar,
    "rating"          float
);

CREATE TABLE IF NOT EXISTS "restaurant_menu_item"
(
    "restaurant_id" bigserial,
    "menu_item_id"  bigserial
);

CREATE TABLE IF NOT EXISTS "menu_item"
(
    "id"           bigserial PRIMARY KEY,
    "name"         varchar,
    "description"  varchar,
    "price"        bigint,
    "decimalprice" varchar,
    "weight"       varchar,
    "adult"        boolean,
    "shippingtype" varchar,
    "available"    boolean
);

ALTER TABLE "restaurant_menu_item"
    ADD FOREIGN KEY ("restaurant_id") REFERENCES "restaurant" ("id");

ALTER TABLE "restaurant_menu_item"
    ADD FOREIGN KEY ("menu_item_id") REFERENCES "menu_item" ("id");
