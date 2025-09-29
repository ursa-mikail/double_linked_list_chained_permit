Linked Permits Demonstrated:

```
CREATE:
create_permit() - Append to end

insert_permit_at_position() - Insert at specific position

create_random_permits() - Bulk creation

READ:
read_permit() - Read by node ID

display_list() - Display all permits

UPDATE:
update_permit_domains() - Modify domains

rotate_key() - Key rotation

revoke_permit() - Soft delete

restore_permit() - Undo revocation

DELETE:
delete_permit() - Delete by node ID

delete_permit_at_position() - Delete by position

Execute:
get_statistics()

```

Out:

```
=== DEMONSTRATING ALL CRUD OPERATIONS ===

1. CREATE OPERATIONS
------------------------------
Creating initial permits...
Created 5 total permits

============================================================
Permit Linked List (Size: 5)
============================================================
Position: 0
Node ID: bfd11a592a67aab7
Domains: ['ai_ml', 'finance']
Key: 0636dc109d972112...
Status: ACTIVE
Timestamp: 4311159.3
----------------------------------------
Position: 1
Node ID: 8b0d02a33192c7e9
Domains: ['finance']
Key: ac565beffe7b037b...
Status: ACTIVE
Timestamp: 4311159.3
----------------------------------------
Position: 2
Node ID: c7a6a3e65531f698
Domains: ['ai_ml']
Key: b2e657e4713f8033...
Status: ACTIVE
Timestamp: 4311159.3
----------------------------------------
Position: 3
Node ID: 7d4871ab1307d55d
Domains: ['marketing', 'development']
Key: dccf33c15536ca77...
Status: ACTIVE
Timestamp: 4311159.3
----------------------------------------
Position: 4
Node ID: ac9c2b856147ea72
Domains: ['ai_ml', 'marketing']
Key: 4c4134e2010f3ec7...
Status: ACTIVE
Timestamp: 4311159.3
----------------------------------------

2. INSERT OPERATIONS
------------------------------
Inserting at position 0 (beginning)...
Inserted permit at position 0 with ID b5b3fcc601181583
Inserting at position 2 (middle)...
Inserted permit at position 2 with ID a61c3d0551f77ecd
Inserting at the end...
Inserted permit at position 7 with ID f383d0d2497c58c8

============================================================
Permit Linked List (Size: 8)
============================================================
Position: 0
Node ID: b5b3fcc601181583
Domains: ['it', 'hr']
Key: dfdac2998c95e127...
Status: ACTIVE
Timestamp: 4311159.3
----------------------------------------
Position: 1
Node ID: bfd11a592a67aab7
Domains: ['ai_ml', 'finance']
Key: 0636dc109d972112...
Status: ACTIVE
Timestamp: 4311159.3
----------------------------------------
Position: 2
Node ID: a61c3d0551f77ecd
Domains: ['marketing']
Key: 675d3e20d08201bb...
Status: ACTIVE
Timestamp: 4311159.3
----------------------------------------
Position: 3
Node ID: 8b0d02a33192c7e9
Domains: ['finance']
Key: ac565beffe7b037b...
Status: ACTIVE
Timestamp: 4311159.3
----------------------------------------
Position: 4
Node ID: c7a6a3e65531f698
Domains: ['ai_ml']
Key: b2e657e4713f8033...
Status: ACTIVE
Timestamp: 4311159.3
----------------------------------------
Position: 5
Node ID: 7d4871ab1307d55d
Domains: ['marketing', 'development']
Key: dccf33c15536ca77...
Status: ACTIVE
Timestamp: 4311159.3
----------------------------------------
Position: 6
Node ID: ac9c2b856147ea72
Domains: ['ai_ml', 'marketing']
Key: 4c4134e2010f3ec7...
Status: ACTIVE
Timestamp: 4311159.3
----------------------------------------
Position: 7
Node ID: f383d0d2497c58c8
Domains: ['legal']
Key: cda7a7902a08bf94...
Status: ACTIVE
Timestamp: 4311159.3
----------------------------------------

3. READ OPERATIONS
------------------------------
Read permit b5b3fcc601181583: ['it', 'hr']
Reading non-existent permit: None

4. UPDATE OPERATIONS
------------------------------
Updated domains for node 8b0d02a33192c7e9: ['finance', 'operations', 'sales']
Key rotated for node c7a6a3e65531f698
Revoked permit with ID 7d4871ab1307d55d

============================================================
Permit Linked List (Size: 8)
============================================================
Position: 0
Node ID: b5b3fcc601181583
Domains: ['it', 'hr']
Key: dfdac2998c95e127...
Status: ACTIVE
Timestamp: 4311159.3
----------------------------------------
Position: 1
Node ID: bfd11a592a67aab7
Domains: ['ai_ml', 'finance']
Key: 0636dc109d972112...
Status: ACTIVE
Timestamp: 4311159.3
----------------------------------------
Position: 2
Node ID: a61c3d0551f77ecd
Domains: ['marketing']
Key: 675d3e20d08201bb...
Status: ACTIVE
Timestamp: 4311159.3
----------------------------------------
Position: 3
Node ID: 8b0d02a33192c7e9
Domains: ['finance', 'operations', 'sales']
Key: ac565beffe7b037b...
Status: ACTIVE
Timestamp: 4311159.3
----------------------------------------
Position: 4
Node ID: c7a6a3e65531f698
Domains: ['ai_ml']
Key: 1754e7dcea793fce...
Status: ACTIVE
Timestamp: 4311159.3
----------------------------------------
Position: 5
Node ID: 7d4871ab1307d55d
Domains: ['marketing', 'development']
Key: dccf33c15536ca77...
Status: REVOKED
Timestamp: 4311159.3
----------------------------------------
Position: 6
Node ID: ac9c2b856147ea72
Domains: ['ai_ml', 'marketing']
Key: 4c4134e2010f3ec7...
Status: ACTIVE
Timestamp: 4311159.3
----------------------------------------
Position: 7
Node ID: f383d0d2497c58c8
Domains: ['legal']
Key: cda7a7902a08bf94...
Status: ACTIVE
Timestamp: 4311159.3
----------------------------------------

5. DELETE OPERATIONS
------------------------------
Deleting by node ID...
Deleted permit with ID a61c3d0551f77ecd
Deleting at position 1...
Deleted permit with ID bfd11a592a67aab7
Attempting to delete non-existent permit...
Node nonexistent not found for deletion
Attempting to delete at invalid position...
Invalid position 100. List size is 6

============================================================
Permit Linked List (Size: 6)
============================================================
Position: 0
Node ID: b5b3fcc601181583
Domains: ['it', 'hr']
Key: dfdac2998c95e127...
Status: ACTIVE
Timestamp: 4311159.3
----------------------------------------
Position: 1
Node ID: 8b0d02a33192c7e9
Domains: ['finance', 'operations', 'sales']
Key: ac565beffe7b037b...
Status: ACTIVE
Timestamp: 4311159.3
----------------------------------------
Position: 2
Node ID: c7a6a3e65531f698
Domains: ['ai_ml']
Key: 1754e7dcea793fce...
Status: ACTIVE
Timestamp: 4311159.3
----------------------------------------
Position: 3
Node ID: 7d4871ab1307d55d
Domains: ['marketing', 'development']
Key: dccf33c15536ca77...
Status: REVOKED
Timestamp: 4311159.3
----------------------------------------
Position: 4
Node ID: ac9c2b856147ea72
Domains: ['ai_ml', 'marketing']
Key: 4c4134e2010f3ec7...
Status: ACTIVE
Timestamp: 4311159.3
----------------------------------------
Position: 5
Node ID: f383d0d2497c58c8
Domains: ['legal']
Key: cda7a7902a08bf94...
Status: ACTIVE
Timestamp: 4311159.3
----------------------------------------

6. RESTORE OPERATION
------------------------------
Restoring revoked permit...
Restored permit with ID 7d4871ab1307d55d

=== FINAL STATISTICS ===
Total permits: 6
Active permits: 6
Revoked permits: 0
Domain distribution: {'it': 1, 'hr': 1, 'finance': 1, 'operations': 1, 'sales': 1, 'ai_ml': 2, 'marketing': 2, 'development': 1, 'legal': 1}
```