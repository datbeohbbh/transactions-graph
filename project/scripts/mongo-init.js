// use db.getSiblingDB to get instance of admin db
db = db.getSiblingDB("graphdb")
// grant role readWrite on database graphdb for user nguyendat1211
db.createUser(
        {
            user: "nguyendat1211",
            pwd: "password",
            roles: [
                {
                    role: "readWrite",
                    db: "graphdb"
                }
            ]
        }
);

db = db.getSiblingDB("graphdb")
db.createUser(
    {
        user: "graph_service",
        pwd: "password",
        roles: [
            {
                role: "read",
                db: "graphdb"
            }
        ]
    }
);

db = db.getSiblingDB("address_manager")
db.createUser(
    {
        user: "address_manager",
        pwd: "password",
        roles: [
            {
                role: "readWrite",
                db: "address_manager"
            }
        ]
    }
);