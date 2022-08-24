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
)

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
)