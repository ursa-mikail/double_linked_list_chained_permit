import os
import secrets
import random
from typing import Optional, Dict, Any, List, Set
from dataclasses import dataclass
from enum import Enum
from faker import Faker

# Initialize Faker
fake = Faker()

class Domain(Enum):
    FINANCE = "finance"
    AI_ML = "ai_ml"
    HR = "hr"
    IT = "it"
    MARKETING = "marketing"
    SALES = "sales"
    OPERATIONS = "operations"
    LEGAL = "legal"
    RESEARCH = "research"
    DEVELOPMENT = "development"

    @classmethod
    def get_random_domain(cls) -> 'Domain':
        """Get a random domain from all available domains"""
        return random.choice(list(cls))
    
    @classmethod
    def get_random_domains(cls, count: int = 1) -> Set['Domain']:
        """Get a set of random unique domains"""
        available_domains = list(cls)
        if count > len(available_domains):
            count = len(available_domains)
        return set(random.sample(available_domains, count))
    
    @classmethod
    def create_fake_departments(cls, count: int) -> List['Domain']:
        """Create random department domains using Faker"""
        department_mapping = {
            'finance': cls.FINANCE,
            'accounting': cls.FINANCE,
            'human resources': cls.HR,
            'hr': cls.HR,
            'information technology': cls.IT,
            'it': cls.IT,
            'marketing': cls.MARKETING,
            'sales': cls.SALES,
            'operations': cls.OPERATIONS,
            'legal': cls.LEGAL,
            'research': cls.RESEARCH,
            'development': cls.DEVELOPMENT,
            'ai': cls.AI_ML,
            'machine learning': cls.AI_ML,
            'data science': cls.AI_ML
        }
        
        departments = []
        for _ in range(count):
            dept_name = fake.word().lower()
            if dept_name in department_mapping:
                departments.append(department_mapping[dept_name])
            else:
                departments.append(cls.get_random_domain())
        
        return departments

@dataclass
class Permit:
    domains: Set[Domain]
    key: bytes  # 16 random bytes
    timestamp: float
    revoked: bool = False

class PermitNode:
    def __init__(self, permit: Permit):
        self.permit = permit
        self.prev: Optional['PermitNode'] = None
        self.next: Optional['PermitNode'] = None
        self.id: str = secrets.token_hex(8)

class PermitLinkedList:
    def __init__(self):
        self.head: Optional[PermitNode] = None
        self.tail: Optional[PermitNode] = None
        self.size = 0
        self.node_map: Dict[str, PermitNode] = {}

    def generate_key(self) -> bytes:
        """Generate 16 random bytes for key"""
        return secrets.token_bytes(16)

    def create_permit(self, domains: Set[Domain]) -> PermitNode:
        """Create a new permit and add to the end of the list"""
        permit = Permit(
            domains=domains,
            key=self.generate_key(),
            timestamp=os.times().elapsed
        )
        
        new_node = PermitNode(permit)
        
        if self.head is None:
            self.head = new_node
            self.tail = new_node
        else:
            new_node.prev = self.tail
            self.tail.next = new_node
            self.tail = new_node
        
        self.size += 1
        self.node_map[new_node.id] = new_node
        return new_node

    def insert_permit_at_position(self, domains: Set[Domain], position: int) -> Optional[PermitNode]:
        """Insert a new permit at a specific position in the list"""
        if position < 0 or position > self.size:
            print(f"Invalid position {position}. List size is {self.size}")
            return None
        
        permit = Permit(
            domains=domains,
            key=self.generate_key(),
            timestamp=os.times().elapsed
        )
        
        new_node = PermitNode(permit)
        
        if position == 0:  # Insert at head
            new_node.next = self.head
            if self.head:
                self.head.prev = new_node
            self.head = new_node
            if self.tail is None:
                self.tail = new_node
                
        elif position == self.size:  # Insert at tail
            new_node.prev = self.tail
            if self.tail:
                self.tail.next = new_node
            self.tail = new_node
            
        else:  # Insert in middle
            current = self.head
            for _ in range(position):
                if current:
                    current = current.next
            
            if current:
                new_node.prev = current.prev
                new_node.next = current
                if current.prev:
                    current.prev.next = new_node
                current.prev = new_node
        
        self.size += 1
        self.node_map[new_node.id] = new_node
        print(f"Inserted permit at position {position} with ID {new_node.id}")
        return new_node

    def create_random_permits(self, count: int = 1) -> List[PermitNode]:
        """Create multiple permits with random domains"""
        nodes = []
        for _ in range(count):
            num_domains = random.randint(1, 3)
            random_domains = Domain.get_random_domains(num_domains)
            node = self.create_permit(random_domains)
            nodes.append(node)
        return nodes

    def create_fake_department_permits(self, count: int = 1) -> List[PermitNode]:
        """Create permits using Faker-generated department names"""
        nodes = []
        for _ in range(count):
            num_domains = random.randint(1, 3)
            fake_departments = Domain.create_fake_departments(num_domains)
            node = self.create_permit(set(fake_departments))
            nodes.append(node)
        return nodes

    def read_permit(self, node_id: str) -> Optional[Permit]:
        """Read a permit by node ID"""
        node = self.node_map.get(node_id)
        return node.permit if node else None

    def update_permit_domains(self, node_id: str, new_domains: Set[Domain]) -> bool:
        """Update domains for a specific permit"""
        node = self.node_map.get(node_id)
        if node and not node.permit.revoked:
            node.permit.domains = new_domains
            node.permit.timestamp = os.times().elapsed
            print(f"Updated domains for node {node_id}: {[d.value for d in new_domains]}")
            return True
        return False

    def delete_permit(self, node_id: str) -> bool:
        """Delete a permit by node ID"""
        node = self.node_map.get(node_id)
        if not node:
            print(f"Node {node_id} not found for deletion")
            return False
        
        # Remove from linked list
        if node.prev:
            node.prev.next = node.next
        else:  # This is the head
            self.head = node.next
        
        if node.next:
            node.next.prev = node.prev
        else:  # This is the tail
            self.tail = node.prev
        
        # Remove from node map
        del self.node_map[node_id]
        self.size -= 1
        print(f"Deleted permit with ID {node_id}")
        return True

    def delete_permit_at_position(self, position: int) -> bool:
        """Delete a permit at a specific position"""
        if position < 0 or position >= self.size:
            print(f"Invalid position {position}. List size is {self.size}")
            return False
        
        current = self.head
        for _ in range(position):
            if current:
                current = current.next
        
        if current:
            return self.delete_permit(current.id)
        return False

    def rotate_key(self, node_id: str) -> bool:
        """Rotate key for a specific permit"""
        node = self.node_map.get(node_id)
        if node and not node.permit.revoked:
            old_key = node.permit.key
            node.permit.key = self.generate_key()
            node.permit.timestamp = os.times().elapsed
            print(f"Key rotated for node {node_id}")
            return True
        return False

    def revoke_permit(self, node_id: str) -> bool:
        """Revoke a permit (soft delete)"""
        node = self.node_map.get(node_id)
        if node and not node.permit.revoked:
            node.permit.revoked = True
            node.permit.timestamp = os.times().elapsed
            print(f"Revoked permit with ID {node_id}")
            return True
        return False

    def restore_permit(self, node_id: str) -> bool:
        """Restore a revoked permit"""
        node = self.node_map.get(node_id)
        if node and node.permit.revoked:
            node.permit.revoked = False
            node.permit.timestamp = os.times().elapsed
            print(f"Restored permit with ID {node_id}")
            return True
        return False

    def bulk_rotate_keys(self, domain: Optional[Domain] = None) -> int:
        """Rotate keys for all permits (optionally filtered by domain)"""
        rotated_count = 0
        current = self.head
        
        while current:
            if not current.permit.revoked and (domain is None or domain in current.permit.domains):
                current.permit.key = self.generate_key()
                current.permit.timestamp = os.times().elapsed
                rotated_count += 1
            current = current.next
        
        return rotated_count

    def find_by_domain(self, domain: Domain) -> List[PermitNode]:
        """Find all permits for a specific domain"""
        result = []
        current = self.head
        
        while current:
            if domain in current.permit.domains and not current.permit.revoked:
                result.append(current)
            current = current.next
        
        return result

    def get_active_permits(self) -> List[PermitNode]:
        """Get all non-revoked permits"""
        result = []
        current = self.head
        
        while current:
            if not current.permit.revoked:
                result.append(current)
            current = current.next
        
        return result

    def display_list(self, show_revoked: bool = False) -> None:
        """Display the entire linked list"""
        current = self.head
        position = 0
        
        print(f"\n{'='*60}")
        print(f"Permit Linked List (Size: {self.size})")
        print(f"{'='*60}")
        
        if self.head is None:
            print("List is empty")
            return
            
        while current:
            status = "REVOKED" if current.permit.revoked else "ACTIVE"
            
            if show_revoked or not current.permit.revoked:
                print(f"Position: {position}")
                print(f"Node ID: {current.id}")
                print(f"Domains: {[domain.value for domain in current.permit.domains]}")
                print(f"Key: {current.permit.key.hex()[:16]}...")  # Show first 16 chars
                print(f"Status: {status}")
                print(f"Timestamp: {current.permit.timestamp}")
                print(f"{'-'*40}")
            
            current = current.next
            position += 1

    def to_list(self) -> List[Dict[str, Any]]:
        """Convert linked list to Python list for serialization"""
        result = []
        current = self.head
        
        while current:
            result.append({
                'node_id': current.id,
                'domains': [domain.value for domain in current.permit.domains],
                'key': current.permit.key.hex(),
                'timestamp': current.permit.timestamp,
                'revoked': current.permit.revoked
            })
            current = current.next
        
        return result

    def get_statistics(self) -> Dict[str, Any]:
        """Get statistics about the permit list"""
        stats = {
            'total_permits': self.size,
            'active_permits': 0,
            'revoked_permits': 0,
            'domain_distribution': {},
            'average_domains_per_permit': 0
        }
        
        current = self.head
        total_domains = 0
        
        while current:
            if current.permit.revoked:
                stats['revoked_permits'] += 1
            else:
                stats['active_permits'] += 1
            
            total_domains += len(current.permit.domains)
            
            for domain in current.permit.domains:
                domain_name = domain.value
                if domain_name not in stats['domain_distribution']:
                    stats['domain_distribution'][domain_name] = 0
                stats['domain_distribution'][domain_name] += 1
            
            current = current.next
        
        if self.size > 0:
            stats['average_domains_per_permit'] = total_domains / self.size
        
        return stats

# Enhanced main function demonstrating ALL CRUD operations
def main():
    # Create permit list
    permit_list = PermitLinkedList()
    
    print("=== DEMONSTRATING ALL CRUD OPERATIONS ===")
    
    # 1. CREATE OPERATIONS
    print("\n1. CREATE OPERATIONS")
    print("-" * 30)
    
    # Create initial permits
    print("Creating initial permits...")
    node1 = permit_list.create_permit({Domain.FINANCE, Domain.AI_ML})
    node2 = permit_list.create_permit({Domain.FINANCE})
    node3 = permit_list.create_permit({Domain.AI_ML})
    
    # Create random permits
    random_nodes = permit_list.create_random_permits(2)
    print(f"Created {3 + len(random_nodes)} total permits")
    
    permit_list.display_list()
    
    # 2. INSERT OPERATIONS (at specific positions)
    print("\n2. INSERT OPERATIONS")
    print("-" * 30)
    
    # Insert at beginning (position 0)
    print("Inserting at position 0 (beginning)...")
    inserted_head = permit_list.insert_permit_at_position({Domain.HR, Domain.IT}, 0)
    
    # Insert at middle (position 2)
    print("Inserting at position 2 (middle)...")
    inserted_middle = permit_list.insert_permit_at_position({Domain.MARKETING}, 2)
    
    # Insert at end (position = size)
    print("Inserting at the end...")
    inserted_tail = permit_list.insert_permit_at_position({Domain.LEGAL}, permit_list.size)
    
    permit_list.display_list()
    
    # 3. READ OPERATIONS
    print("\n3. READ OPERATIONS")
    print("-" * 30)
    
    if inserted_head:
        permit_data = permit_list.read_permit(inserted_head.id)
        if permit_data:
            print(f"Read permit {inserted_head.id}: {[d.value for d in permit_data.domains]}")
    
    # Try to read non-existent permit
    non_existent = permit_list.read_permit("nonexistent")
    print(f"Reading non-existent permit: {non_existent}")
    
    # 4. UPDATE OPERATIONS
    print("\n4. UPDATE OPERATIONS")
    print("-" * 30)
    
    # Update domains
    if node2:
        permit_list.update_permit_domains(node2.id, {Domain.FINANCE, Domain.SALES, Domain.OPERATIONS})
    
    # Rotate key
    if node3:
        permit_list.rotate_key(node3.id)
    
    # Revoke a permit
    if random_nodes:
        permit_list.revoke_permit(random_nodes[0].id)
    
    permit_list.display_list(show_revoked=True)
    
    # 5. DELETE OPERATIONS
    print("\n5. DELETE OPERATIONS")
    print("-" * 30)
    
    # Delete by ID
    if inserted_middle:
        print("Deleting by node ID...")
        permit_list.delete_permit(inserted_middle.id)
    
    # Delete by position
    print("Deleting at position 1...")
    permit_list.delete_permit_at_position(1)
    
    # Try to delete non-existent permit
    print("Attempting to delete non-existent permit...")
    permit_list.delete_permit("nonexistent")
    
    # Try to delete at invalid position
    print("Attempting to delete at invalid position...")
    permit_list.delete_permit_at_position(100)
    
    permit_list.display_list(show_revoked=True)
    
    # 6. RESTORE OPERATION (undo revocation)
    print("\n6. RESTORE OPERATION")
    print("-" * 30)
    
    if random_nodes and random_nodes[0]:
        print("Restoring revoked permit...")
        permit_list.restore_permit(random_nodes[0].id)
    
    final_stats = permit_list.get_statistics()
    print(f"\n=== FINAL STATISTICS ===")
    print(f"Total permits: {final_stats['total_permits']}")
    print(f"Active permits: {final_stats['active_permits']}")
    print(f"Revoked permits: {final_stats['revoked_permits']}")
    print(f"Domain distribution: {final_stats['domain_distribution']}")

if __name__ == "__main__":
    main()

    