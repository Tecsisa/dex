Groups Design
===============================

## Diagram

```
                                  Client
                                    |
                                    |
                                    |
                    +-----+         |
                    |     |         |
          +---------v--+  |         |     +------------------+      +-------------+
          |            +--+         |     |                  |      |             |
          |            |          +-------v--------+         |      |             |
          | SuperAdmin <----------+                |        +-------v--+       +-------+
          |            |          |  Organization  +--------+  Owners  +-------+ Users |
          |            +----------+                |        +----------+       +-------+
          +------------+          +----------------+


          +------------->  Admin                 Sample:
                                                     Organization: tecsisa
          +-------------+  Membership                Owners: admin_tecsisa
                                                     Organization/tecsisa/admin_tecsisa/users
          +------+
          |      |  Group
          +------+

```
