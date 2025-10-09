package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/jaswdr/faker"
)

// Domain represents different business domains
type Domain string

const (
	DomainFinance     Domain = "finance"
	DomainAIML        Domain = "ai_ml"
	DomainHR          Domain = "hr"
	DomainIT          Domain = "it"
	DomainMarketing   Domain = "marketing"
	DomainSales       Domain = "sales"
	DomainOperations  Domain = "operations"
	DomainLegal       Domain = "legal"
	DomainResearch    Domain = "research"
	DomainDevelopment Domain = "development"
)

var allDomains = []Domain{
	DomainFinance, DomainAIML, DomainHR, DomainIT, DomainMarketing,
	DomainSales, DomainOperations, DomainLegal, DomainResearch, DomainDevelopment,
}

// GetRandomDomain returns a random domain
func GetRandomDomain() Domain {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(allDomains))))
	return allDomains[n.Int64()]
}

// GetRandomDomains returns a set of unique random domains
func GetRandomDomains(count int) map[Domain]bool {
	if count > len(allDomains) {
		count = len(allDomains)
	}
	
	domains := make(map[Domain]bool)
	for len(domains) < count {
		domains[GetRandomDomain()] = true
	}
	return domains
}

// CreateFakeDepartments creates random department domains using Faker
func CreateFakeDepartments(count int) []Domain {
	fake := faker.New()
	departmentMapping := map[string]Domain{
		"finance":              DomainFinance,
		"accounting":           DomainFinance,
		"human resources":      DomainHR,
		"hr":                   DomainHR,
		"information technology": DomainIT,
		"it":                   DomainIT,
		"marketing":            DomainMarketing,
		"sales":                DomainSales,
		"operations":           DomainOperations,
		"legal":                DomainLegal,
		"research":             DomainResearch,
		"development":          DomainDevelopment,
		"ai":                   DomainAIML,
		"machine learning":     DomainAIML,
		"data science":         DomainAIML,
	}
	
	departments := make([]Domain, count)
	for i := 0; i < count; i++ {
		deptName := strings.ToLower(fake.Lorem().Word())
		if domain, ok := departmentMapping[deptName]; ok {
			departments[i] = domain
		} else {
			departments[i] = GetRandomDomain()
		}
	}
	
	return departments
}

// Permit represents a permission with domains and key
type Permit struct {
	Domains   map[Domain]bool
	Key       []byte
	Timestamp float64
	Revoked   bool
}

// PermitNode is a node in the doubly linked list
type PermitNode struct {
	Permit *Permit
	Prev   *PermitNode
	Next   *PermitNode
	ID     string
}

// PermitLinkedList manages permits in a doubly linked list
type PermitLinkedList struct {
	Head    *PermitNode
	Tail    *PermitNode
	Size    int
	NodeMap map[string]*PermitNode
}

// NewPermitLinkedList creates a new permit linked list
func NewPermitLinkedList() *PermitLinkedList {
	return &PermitLinkedList{
		NodeMap: make(map[string]*PermitNode),
	}
}

// GenerateKey generates 16 random bytes for key
func (pll *PermitLinkedList) GenerateKey() []byte {
	key := make([]byte, 16)
	rand.Read(key)
	return key
}

// GenerateID generates a random node ID
func generateID() string {
	bytes := make([]byte, 8)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// getCurrentTimestamp returns current time as float64 seconds
func getCurrentTimestamp() float64 {
	return float64(time.Now().UnixNano()) / 1e9
}

// CreatePermit creates a new permit and adds to the end of the list
func (pll *PermitLinkedList) CreatePermit(domains map[Domain]bool) *PermitNode {
	permit := &Permit{
		Domains:   domains,
		Key:       pll.GenerateKey(),
		Timestamp: getCurrentTimestamp(),
		Revoked:   false,
	}
	
	newNode := &PermitNode{
		Permit: permit,
		ID:     generateID(),
	}
	
	if pll.Head == nil {
		pll.Head = newNode
		pll.Tail = newNode
	} else {
		newNode.Prev = pll.Tail
		pll.Tail.Next = newNode
		pll.Tail = newNode
	}
	
	pll.Size++
	pll.NodeMap[newNode.ID] = newNode
	return newNode
}

// InsertPermitAtPosition inserts a new permit at a specific position
func (pll *PermitLinkedList) InsertPermitAtPosition(domains map[Domain]bool, position int) *PermitNode {
	if position < 0 || position > pll.Size {
		fmt.Printf("Invalid position %d. List size is %d\n", position, pll.Size)
		return nil
	}
	
	permit := &Permit{
		Domains:   domains,
		Key:       pll.GenerateKey(),
		Timestamp: getCurrentTimestamp(),
		Revoked:   false,
	}
	
	newNode := &PermitNode{
		Permit: permit,
		ID:     generateID(),
	}
	
	if position == 0 { // Insert at head
		newNode.Next = pll.Head
		if pll.Head != nil {
			pll.Head.Prev = newNode
		}
		pll.Head = newNode
		if pll.Tail == nil {
			pll.Tail = newNode
		}
	} else if position == pll.Size { // Insert at tail
		newNode.Prev = pll.Tail
		if pll.Tail != nil {
			pll.Tail.Next = newNode
		}
		pll.Tail = newNode
	} else { // Insert in middle
		current := pll.Head
		for i := 0; i < position; i++ {
			if current != nil {
				current = current.Next
			}
		}
		
		if current != nil {
			newNode.Prev = current.Prev
			newNode.Next = current
			if current.Prev != nil {
				current.Prev.Next = newNode
			}
			current.Prev = newNode
		}
	}
	
	pll.Size++
	pll.NodeMap[newNode.ID] = newNode
	fmt.Printf("Inserted permit at position %d with ID %s\n", position, newNode.ID)
	return newNode
}

// CreateRandomPermits creates multiple permits with random domains
func (pll *PermitLinkedList) CreateRandomPermits(count int) []*PermitNode {
	nodes := make([]*PermitNode, count)
	for i := 0; i < count; i++ {
		numDomains, _ := rand.Int(rand.Reader, big.NewInt(3))
		randomDomains := GetRandomDomains(int(numDomains.Int64()) + 1)
		nodes[i] = pll.CreatePermit(randomDomains)
	}
	return nodes
}

// CreateFakeDepartmentPermits creates permits using Faker-generated department names
func (pll *PermitLinkedList) CreateFakeDepartmentPermits(count int) []*PermitNode {
	nodes := make([]*PermitNode, count)
	for i := 0; i < count; i++ {
		numDomains, _ := rand.Int(rand.Reader, big.NewInt(3))
		fakeDepartments := CreateFakeDepartments(int(numDomains.Int64()) + 1)
		domainSet := make(map[Domain]bool)
		for _, dept := range fakeDepartments {
			domainSet[dept] = true
		}
		nodes[i] = pll.CreatePermit(domainSet)
	}
	return nodes
}

// ReadPermit reads a permit by node ID
func (pll *PermitLinkedList) ReadPermit(nodeID string) *Permit {
	if node, ok := pll.NodeMap[nodeID]; ok {
		return node.Permit
	}
	return nil
}

// UpdatePermitDomains updates domains for a specific permit
func (pll *PermitLinkedList) UpdatePermitDomains(nodeID string, newDomains map[Domain]bool) bool {
	node, ok := pll.NodeMap[nodeID]
	if ok && !node.Permit.Revoked {
		node.Permit.Domains = newDomains
		node.Permit.Timestamp = getCurrentTimestamp()
		domainNames := make([]string, 0, len(newDomains))
		for d := range newDomains {
			domainNames = append(domainNames, string(d))
		}
		fmt.Printf("Updated domains for node %s: %v\n", nodeID, domainNames)
		return true
	}
	return false
}

// DeletePermit deletes a permit by node ID
func (pll *PermitLinkedList) DeletePermit(nodeID string) bool {
	node, ok := pll.NodeMap[nodeID]
	if !ok {
		fmt.Printf("Node %s not found for deletion\n", nodeID)
		return false
	}
	
	// Remove from linked list
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else { // This is the head
		pll.Head = node.Next
	}
	
	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else { // This is the tail
		pll.Tail = node.Prev
	}
	
	// Remove from node map
	delete(pll.NodeMap, nodeID)
	pll.Size--
	fmt.Printf("Deleted permit with ID %s\n", nodeID)
	return true
}

// DeletePermitAtPosition deletes a permit at a specific position
func (pll *PermitLinkedList) DeletePermitAtPosition(position int) bool {
	if position < 0 || position >= pll.Size {
		fmt.Printf("Invalid position %d. List size is %d\n", position, pll.Size)
		return false
	}
	
	current := pll.Head
	for i := 0; i < position; i++ {
		if current != nil {
			current = current.Next
		}
	}
	
	if current != nil {
		return pll.DeletePermit(current.ID)
	}
	return false
}

// RotateKey rotates key for a specific permit
func (pll *PermitLinkedList) RotateKey(nodeID string) bool {
	node, ok := pll.NodeMap[nodeID]
	if ok && !node.Permit.Revoked {
		node.Permit.Key = pll.GenerateKey()
		node.Permit.Timestamp = getCurrentTimestamp()
		fmt.Printf("Key rotated for node %s\n", nodeID)
		return true
	}
	return false
}

// RevokePermit revokes a permit (soft delete)
func (pll *PermitLinkedList) RevokePermit(nodeID string) bool {
	node, ok := pll.NodeMap[nodeID]
	if ok && !node.Permit.Revoked {
		node.Permit.Revoked = true
		node.Permit.Timestamp = getCurrentTimestamp()
		fmt.Printf("Revoked permit with ID %s\n", nodeID)
		return true
	}
	return false
}

// RestorePermit restores a revoked permit
func (pll *PermitLinkedList) RestorePermit(nodeID string) bool {
	node, ok := pll.NodeMap[nodeID]
	if ok && node.Permit.Revoked {
		node.Permit.Revoked = false
		node.Permit.Timestamp = getCurrentTimestamp()
		fmt.Printf("Restored permit with ID %s\n", nodeID)
		return true
	}
	return false
}

// BulkRotateKeys rotates keys for all permits (optionally filtered by domain)
func (pll *PermitLinkedList) BulkRotateKeys(domain *Domain) int {
	rotatedCount := 0
	current := pll.Head
	
	for current != nil {
		if !current.Permit.Revoked && (domain == nil || current.Permit.Domains[*domain]) {
			current.Permit.Key = pll.GenerateKey()
			current.Permit.Timestamp = getCurrentTimestamp()
			rotatedCount++
		}
		current = current.Next
	}
	
	return rotatedCount
}

// FindByDomain finds all permits for a specific domain
func (pll *PermitLinkedList) FindByDomain(domain Domain) []*PermitNode {
	result := []*PermitNode{}
	current := pll.Head
	
	for current != nil {
		if current.Permit.Domains[domain] && !current.Permit.Revoked {
			result = append(result, current)
		}
		current = current.Next
	}
	
	return result
}

// GetActivePermits gets all non-revoked permits
func (pll *PermitLinkedList) GetActivePermits() []*PermitNode {
	result := []*PermitNode{}
	current := pll.Head
	
	for current != nil {
		if !current.Permit.Revoked {
			result = append(result, current)
		}
		current = current.Next
	}
	
	return result
}

// DisplayList displays the entire linked list
func (pll *PermitLinkedList) DisplayList(showRevoked bool) {
	current := pll.Head
	position := 0
	
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Printf("Permit Linked List (Size: %d)\n", pll.Size)
	fmt.Println(strings.Repeat("=", 60))
	
	if pll.Head == nil {
		fmt.Println("List is empty")
		return
	}
	
	for current != nil {
		status := "ACTIVE"
		if current.Permit.Revoked {
			status = "REVOKED"
		}
		
		if showRevoked || !current.Permit.Revoked {
			fmt.Printf("Position: %d\n", position)
			fmt.Printf("Node ID: %s\n", current.ID)
			
			domainNames := make([]string, 0, len(current.Permit.Domains))
			for d := range current.Permit.Domains {
				domainNames = append(domainNames, string(d))
			}
			fmt.Printf("Domains: %v\n", domainNames)
			
			keyHex := hex.EncodeToString(current.Permit.Key)
			if len(keyHex) > 16 {
				keyHex = keyHex[:16] + "..."
			}
			fmt.Printf("Key: %s\n", keyHex)
			fmt.Printf("Status: %s\n", status)
			fmt.Printf("Timestamp: %.6f\n", current.Permit.Timestamp)
			fmt.Println(strings.Repeat("-", 40))
		}
		
		current = current.Next
		position++
	}
}

// Statistics represents permit list statistics
type Statistics struct {
	TotalPermits            int
	ActivePermits           int
	RevokedPermits          int
	DomainDistribution      map[string]int
	AverageDomainsPerPermit float64
}

// GetStatistics gets statistics about the permit list
func (pll *PermitLinkedList) GetStatistics() Statistics {
	stats := Statistics{
		TotalPermits:       pll.Size,
		DomainDistribution: make(map[string]int),
	}
	
	current := pll.Head
	totalDomains := 0
	
	for current != nil {
		if current.Permit.Revoked {
			stats.RevokedPermits++
		} else {
			stats.ActivePermits++
		}
		
		totalDomains += len(current.Permit.Domains)
		
		for domain := range current.Permit.Domains {
			stats.DomainDistribution[string(domain)]++
		}
		
		current = current.Next
	}
	
	if pll.Size > 0 {
		stats.AverageDomainsPerPermit = float64(totalDomains) / float64(pll.Size)
	}
	
	return stats
}

func main() {
	// Create permit list
	permitList := NewPermitLinkedList()
	
	fmt.Println("=== DEMONSTRATING ALL CRUD OPERATIONS ===")
	
	// 1. CREATE OPERATIONS
	fmt.Println("\n1. CREATE OPERATIONS")
	fmt.Println(strings.Repeat("-", 30))
	
	// Create initial permits
	fmt.Println("Creating initial permits...")
	_ = permitList.CreatePermit(map[Domain]bool{DomainFinance: true, DomainAIML: true})
	node2 := permitList.CreatePermit(map[Domain]bool{DomainFinance: true})
	node3 := permitList.CreatePermit(map[Domain]bool{DomainAIML: true})
	
	// Create random permits
	randomNodes := permitList.CreateRandomPermits(2)
	fmt.Printf("Created %d total permits\n", 3+len(randomNodes))
	
	permitList.DisplayList(false)
	
	// 2. INSERT OPERATIONS (at specific positions)
	fmt.Println("\n2. INSERT OPERATIONS")
	fmt.Println(strings.Repeat("-", 30))
	
	// Insert at beginning (position 0)
	fmt.Println("Inserting at position 0 (beginning)...")
	insertedHead := permitList.InsertPermitAtPosition(map[Domain]bool{DomainHR: true, DomainIT: true}, 0)
	
	// Insert at middle (position 2)
	fmt.Println("Inserting at position 2 (middle)...")
	insertedMiddle := permitList.InsertPermitAtPosition(map[Domain]bool{DomainMarketing: true}, 2)
	
	// Insert at end (position = size)
	fmt.Println("Inserting at the end...")
	_ = permitList.InsertPermitAtPosition(map[Domain]bool{DomainLegal: true}, permitList.Size)
	
	permitList.DisplayList(false)
	
	// 3. READ OPERATIONS
	fmt.Println("\n3. READ OPERATIONS")
	fmt.Println(strings.Repeat("-", 30))
	
	if insertedHead != nil {
		permitData := permitList.ReadPermit(insertedHead.ID)
		if permitData != nil {
			domainNames := make([]string, 0, len(permitData.Domains))
			for d := range permitData.Domains {
				domainNames = append(domainNames, string(d))
			}
			fmt.Printf("Read permit %s: %v\n", insertedHead.ID, domainNames)
		}
	}
	
	// Try to read non-existent permit
	nonExistent := permitList.ReadPermit("nonexistent")
	fmt.Printf("Reading non-existent permit: %v\n", nonExistent)
	
	// 4. UPDATE OPERATIONS
	fmt.Println("\n4. UPDATE OPERATIONS")
	fmt.Println(strings.Repeat("-", 30))
	
	// Update domains
	if node2 != nil {
		permitList.UpdatePermitDomains(node2.ID, map[Domain]bool{
			DomainFinance:    true,
			DomainSales:      true,
			DomainOperations: true,
		})
	}
	
	// Rotate key
	if node3 != nil {
		permitList.RotateKey(node3.ID)
	}
	
	// Revoke a permit
	if len(randomNodes) > 0 {
		permitList.RevokePermit(randomNodes[0].ID)
	}
	
	permitList.DisplayList(true)
	
	// 5. DELETE OPERATIONS
	fmt.Println("\n5. DELETE OPERATIONS")
	fmt.Println(strings.Repeat("-", 30))
	
	// Delete by ID
	if insertedMiddle != nil {
		fmt.Println("Deleting by node ID...")
		permitList.DeletePermit(insertedMiddle.ID)
	}
	
	// Delete by position
	fmt.Println("Deleting at position 1...")
	permitList.DeletePermitAtPosition(1)
	
	// Try to delete non-existent permit
	fmt.Println("Attempting to delete non-existent permit...")
	permitList.DeletePermit("nonexistent")
	
	// Try to delete at invalid position
	fmt.Println("Attempting to delete at invalid position...")
	permitList.DeletePermitAtPosition(100)
	
	permitList.DisplayList(true)
	
	// 6. RESTORE OPERATION (undo revocation)
	fmt.Println("\n6. RESTORE OPERATION")
	fmt.Println(strings.Repeat("-", 30))
	
	if len(randomNodes) > 0 && randomNodes[0] != nil {
		fmt.Println("Restoring revoked permit...")
		permitList.RestorePermit(randomNodes[0].ID)
	}
	
	finalStats := permitList.GetStatistics()
	fmt.Println("\n=== FINAL STATISTICS ===")
	fmt.Printf("Total permits: %d\n", finalStats.TotalPermits)
	fmt.Printf("Active permits: %d\n", finalStats.ActivePermits)
	fmt.Printf("Revoked permits: %d\n", finalStats.RevokedPermits)
	fmt.Printf("Domain distribution: %v\n", finalStats.DomainDistribution)
}

/*
# Initialize Go module
go mod init permit-system

# Install the Faker dependency
go get github.com/jaswdr/faker

# Run the program
go run main.go

=== DEMONSTRATING ALL CRUD OPERATIONS ===

1. CREATE OPERATIONS
------------------------------
Creating initial permits...
Created 5 total permits

============================================================
Permit Linked List (Size: 5)
============================================================
Position: 0
Node ID: 22655c0d85cddf89
Domains: [finance ai_ml]
Key: d4d52aa9b10cceae...
Status: ACTIVE
Timestamp: 1760038323.742507
----------------------------------------
Position: 1
Node ID: c5a03539ac5d8ac5
Domains: [finance]
Key: 7d3b27500bf86433...
Status: ACTIVE
Timestamp: 1760038323.742507
----------------------------------------
Position: 2
Node ID: 7039fa0d3784da16
Domains: [ai_ml]
Key: ac4513b2e67a8071...
Status: ACTIVE
Timestamp: 1760038323.742508
----------------------------------------
Position: 3
Node ID: 568e5bdd7cc67dc1
Domains: [operations ai_ml]
Key: f287335c24c28575...
Status: ACTIVE
Timestamp: 1760038323.742521
----------------------------------------
Position: 4
Node ID: 2cbc7da0f2537540
Domains: [legal development finance]
Key: 2f0e07a9acd39bed...
Status: ACTIVE
Timestamp: 1760038323.742523
----------------------------------------

2. INSERT OPERATIONS
------------------------------
Inserting at position 0 (beginning)...
Inserted permit at position 0 with ID 83daf526b0572312
Inserting at position 2 (middle)...
Inserted permit at position 2 with ID c6e8a763fdc57ff0
Inserting at the end...
Inserted permit at position 7 with ID 7d974a87d7f15842

============================================================
Permit Linked List (Size: 8)
============================================================
Position: 0
Node ID: 83daf526b0572312
Domains: [hr it]
Key: 87cd688f1fe40e56...
Status: ACTIVE
Timestamp: 1760038323.742670
----------------------------------------
Position: 1
Node ID: 22655c0d85cddf89
Domains: [ai_ml finance]
Key: d4d52aa9b10cceae...
Status: ACTIVE
Timestamp: 1760038323.742507
----------------------------------------
Position: 2
Node ID: c6e8a763fdc57ff0
Domains: [marketing]
Key: f43925ed55130c60...
Status: ACTIVE
Timestamp: 1760038323.742688
----------------------------------------
Position: 3
Node ID: c5a03539ac5d8ac5
Domains: [finance]
Key: 7d3b27500bf86433...
Status: ACTIVE
Timestamp: 1760038323.742507
----------------------------------------
Position: 4
Node ID: 7039fa0d3784da16
Domains: [ai_ml]
Key: ac4513b2e67a8071...
Status: ACTIVE
Timestamp: 1760038323.742508
----------------------------------------
Position: 5
Node ID: 568e5bdd7cc67dc1
Domains: [operations ai_ml]
Key: f287335c24c28575...
Status: ACTIVE
Timestamp: 1760038323.742521
----------------------------------------
Position: 6
Node ID: 2cbc7da0f2537540
Domains: [legal development finance]
Key: 2f0e07a9acd39bed...
Status: ACTIVE
Timestamp: 1760038323.742523
----------------------------------------
Position: 7
Node ID: 7d974a87d7f15842
Domains: [legal]
Key: 9cc9bc3c8f6cf32d...
Status: ACTIVE
Timestamp: 1760038323.742692
----------------------------------------

3. READ OPERATIONS
------------------------------
Read permit 83daf526b0572312: [it hr]
Reading non-existent permit: <nil>

4. UPDATE OPERATIONS
------------------------------
Updated domains for node c5a03539ac5d8ac5: [sales operations finance]
Key rotated for node 7039fa0d3784da16
Revoked permit with ID 568e5bdd7cc67dc1

============================================================
Permit Linked List (Size: 8)
============================================================
Position: 0
Node ID: 83daf526b0572312
Domains: [hr it]
Key: 87cd688f1fe40e56...
Status: ACTIVE
Timestamp: 1760038323.742670
----------------------------------------
Position: 1
Node ID: 22655c0d85cddf89
Domains: [ai_ml finance]
Key: d4d52aa9b10cceae...
Status: ACTIVE
Timestamp: 1760038323.742507
----------------------------------------
Position: 2
Node ID: c6e8a763fdc57ff0
Domains: [marketing]
Key: f43925ed55130c60...
Status: ACTIVE
Timestamp: 1760038323.742688
----------------------------------------
Position: 3
Node ID: c5a03539ac5d8ac5
Domains: [finance sales operations]
Key: 7d3b27500bf86433...
Status: ACTIVE
Timestamp: 1760038323.742746
----------------------------------------
Position: 4
Node ID: 7039fa0d3784da16
Domains: [ai_ml]
Key: aef23b6f179d43dc...
Status: ACTIVE
Timestamp: 1760038323.742748
----------------------------------------
Position: 5
Node ID: 568e5bdd7cc67dc1
Domains: [operations ai_ml]
Key: f287335c24c28575...
Status: REVOKED
Timestamp: 1760038323.742748
----------------------------------------
Position: 6
Node ID: 2cbc7da0f2537540
Domains: [legal development finance]
Key: 2f0e07a9acd39bed...
Status: ACTIVE
Timestamp: 1760038323.742523
----------------------------------------
Position: 7
Node ID: 7d974a87d7f15842
Domains: [legal]
Key: 9cc9bc3c8f6cf32d...
Status: ACTIVE
Timestamp: 1760038323.742692
----------------------------------------

5. DELETE OPERATIONS
------------------------------
Deleting by node ID...
Deleted permit with ID c6e8a763fdc57ff0
Deleting at position 1...
Deleted permit with ID 22655c0d85cddf89
Attempting to delete non-existent permit...
Node nonexistent not found for deletion
Attempting to delete at invalid position...
Invalid position 100. List size is 6

============================================================
Permit Linked List (Size: 6)
============================================================
Position: 0
Node ID: 83daf526b0572312
Domains: [hr it]
Key: 87cd688f1fe40e56...
Status: ACTIVE
Timestamp: 1760038323.742670
----------------------------------------
Position: 1
Node ID: c5a03539ac5d8ac5
Domains: [sales operations finance]
Key: 7d3b27500bf86433...
Status: ACTIVE
Timestamp: 1760038323.742746
----------------------------------------
Position: 2
Node ID: 7039fa0d3784da16
Domains: [ai_ml]
Key: aef23b6f179d43dc...
Status: ACTIVE
Timestamp: 1760038323.742748
----------------------------------------
Position: 3
Node ID: 568e5bdd7cc67dc1
Domains: [operations ai_ml]
Key: f287335c24c28575...
Status: REVOKED
Timestamp: 1760038323.742748
----------------------------------------
Position: 4
Node ID: 2cbc7da0f2537540
Domains: [development finance legal]
Key: 2f0e07a9acd39bed...
Status: ACTIVE
Timestamp: 1760038323.742523
----------------------------------------
Position: 5
Node ID: 7d974a87d7f15842
Domains: [legal]
Key: 9cc9bc3c8f6cf32d...
Status: ACTIVE
Timestamp: 1760038323.742692
----------------------------------------

6. RESTORE OPERATION
------------------------------
Restoring revoked permit...
Restored permit with ID 568e5bdd7cc67dc1

=== FINAL STATISTICS ===
Total permits: 6
Active permits: 6
Revoked permits: 0
Domain distribution: map[ai_ml:2 development:1 finance:2 hr:1 it:1 legal:2 operations:2 sales:1]

*/