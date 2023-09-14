## SQL Queries to creaate the database

CREATE TABLE customer (
    customer_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    email text,
    name text
);

CREATE TABLE breeder (
    breeder_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    kennel_name text,
    email text,
    cnpj text,
    document text,
    logo text
)

CREATE TABLE sales (
    sale_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    father_id uuid references sales(sale_id),
    mother_id uuid references sales(sale_id),
    breeder_id uuid references breeder(breeder_id),
    is_litter boolean,
    litter_expected_birth_date date,
    litter_expected_amount integer,
    litter_confirmed_amount integer,
    breed text,
    price float,
    shipping_age integer,
    birth_date date,
    shipping text,
    vaccines text[],
    microchip boolean,
    pedigree text,
    weight float,
    height float,
    color text,
    gender char(1),
    traits text[],
    adult_max_height float,
    adult_max_weight float,
    adult_min_height float,
    adult_min_weight float,
    images text[]
);

CREATE TABLE priority (
    priority_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    sale_id uuid references sales(sale_id),
    customer_id uuid references customer(customer_id),
    position integer,
    price float,
    is_available boolean
);
