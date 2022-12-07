## General
L'idée de base a été de dissocier la requête vers la web-app de l'enregistrement en base de données dans le but de lisser d'éventuels bursts.

Pour cela, dès reception de la requête nous allons envoyer un message dans rabbitmq contenant le payload de la requête sans effectuer la moindre validation.

Parrallèlement, les workers vont consommer ces messages, les valider puis les enregistrer dans une base de données

Les messages considérés comme non valides seront envoyés dans une dead letter queue si celle ci existe

## Models
J'ai utilisé 2 tables tickets et products au format suivant:
```
+------------+-----------------------------+--------------------------------------+
| Column     | Type                        | Modifiers                            |
|------------+-----------------------------+--------------------------------------|
| uuid       | uuid                        |  not null default uuid_generate_v4() |
| created_at | timestamp without time zone |  not null                            |
| updated_at | timestamp without time zone |  not null                            |
| order_id   | integer                     |  not null                            |
| vat        | numeric                     |  not null                            |
| total      | numeric                     |  not null                            |
+------------+-----------------------------+--------------------------------------+
Indexes:
    "tickets_pkey" PRIMARY KEY, btree (uuid)
Referenced by:
    TABLE "products" CONSTRAINT "fk_ticket_uuid" FOREIGN KEY (ticket_uuid) REFERENCES tickets(uuid) ON DELETE CASCADE DEFERRABLE
```

```
+-------------+-----------------------------+--------------------------------------+
| Column      | Type                        | Modifiers                            |
|-------------+-----------------------------+--------------------------------------|
| uuid        | uuid                        |  not null default uuid_generate_v4() |
| created_at  | timestamp without time zone |  not null                            |
| updated_at  | timestamp without time zone |  not null                            |
| product_id  | character varying(64)       |  not null                            |
| name        | character varying(128)      |  not null                            |
| price       | numeric                     |  not null                            |
| ticket_uuid | uuid                        |  not null                            |
+-------------+-----------------------------+--------------------------------------+
Indexes:
    "products_pkey" PRIMARY KEY, btree (uuid)
Foreign-key constraints:
    "fk_ticket_uuid" FOREIGN KEY (ticket_uuid) REFERENCES tickets(uuid) ON DELETE CASCADE DEFERRABLE

```

Hormis les contraintes `not null` je n'ai fait aucune supposition concernant d'éventuelles contraintes notamment d'unicité.

Si les champs "product_id" et "order_id" sont bien uniques pour tous les utilisateurs, alors on aurait pu les utiliser comme clé primaire de chaque table et passer à une relation many to many.

De même, je n'ai ajouté aucun index ne sachant pas quelles requêtes seront effectuées sur ces tables.

## Launch
- Créer un fichier .env  à la racine du projet contenant les variables d'environnement décrites dans .env.conf
- Créer et lancer la base de données
- Créer un exchange/binding/queue (+dlq?) et lancer rabbitmq
- `go run main.go -r webApp`
- `go run main.go -r webWorkers`