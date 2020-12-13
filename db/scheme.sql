CREATE TABLE "machines" (
    "id" serial   NOT NULL,
    "name" varchar(255)   NOT NULL,
    "cpu" int   NOT NULL,
    CONSTRAINT "pk_machines" PRIMARY KEY (
        "id"
     )
);

CREATE TABLE "disks" (
    "id" serial   NOT NULL,
    "space" int   NOT NULL,
    CONSTRAINT "pk_disks" PRIMARY KEY (
        "id"
     )
);

CREATE TABLE "machines_disks" (
    "machineId" int   NOT NULL,
    "diskId" int   NOT NULL
);

ALTER TABLE "machines_disks" ADD CONSTRAINT "fk_machines_disks_machineId" FOREIGN KEY("machineId")
REFERENCES "machines" ("id");

ALTER TABLE "machines_disks" ADD CONSTRAINT "fk_machines_disks_diskId" FOREIGN KEY("diskId")
REFERENCES "disks" ("id");

CREATE INDEX "idx_machines_name"
ON "machines" ("name");

INSERT INTO machines ("name", cpu) VALUES ('server-1', 2);
INSERT INTO machines ("name", cpu) VALUES ('server-2', 4);
INSERT INTO machines ("name", cpu) VALUES ('server-3', 6);
INSERT INTO machines ("name", cpu) VALUES ('server-4', 8);
INSERT INTO machines ("name", cpu) VALUES ('server-5', 16);

INSERT INTO disks (space) VALUES (30769108);    
INSERT INTO disks (space) VALUES (2274250);     
INSERT INTO disks (space) VALUES (25506140);    
INSERT INTO disks (space) VALUES (19145716);    
INSERT INTO disks (space) VALUES (21708020);    
INSERT INTO disks (space) VALUES (8346269);     
INSERT INTO disks (space) VALUES (3200380);     
INSERT INTO disks (space) VALUES (21219411);    
INSERT INTO disks (space) VALUES (7392728);     
INSERT INTO disks (space) VALUES (27945070);    
INSERT INTO disks (space) VALUES (5774035);     
INSERT INTO disks (space) VALUES (18171344);    
INSERT INTO disks (space) VALUES (28342767);    
INSERT INTO disks (space) VALUES (4035147);     
INSERT INTO disks (space) VALUES (25111662);    

INSERT INTO machines_disks ("machineId", "diskId") VALUES (1, 1);
INSERT INTO machines_disks ("machineId", "diskId") VALUES (2, 2);
INSERT INTO machines_disks ("machineId", "diskId") VALUES (2, 3);
INSERT INTO machines_disks ("machineId", "diskId") VALUES (2, 4);
INSERT INTO machines_disks ("machineId", "diskId") VALUES (3, 5);
INSERT INTO machines_disks ("machineId", "diskId") VALUES (3, 6);
INSERT INTO machines_disks ("machineId", "diskId") VALUES (4, 7);
INSERT INTO machines_disks ("machineId", "diskId") VALUES (4, 8);
INSERT INTO machines_disks ("machineId", "diskId") VALUES (4, 9);
INSERT INTO machines_disks ("machineId", "diskId") VALUES (4, 10);
INSERT INTO machines_disks ("machineId", "diskId") VALUES (4, 11);
INSERT INTO machines_disks ("machineId", "diskId") VALUES (4, 12);
INSERT INTO machines_disks ("machineId", "diskId") VALUES (5, 13);
INSERT INTO machines_disks ("machineId", "diskId") VALUES (5, 14);
INSERT INTO machines_disks ("machineId", "diskId") VALUES (5, 15);
