# Linked List

## What is a Linked List?

A linked list is a linear data structure, in which the elements are not stored at contiguous memory locations. The elements in a linked list are called nodes. Each node contains a data field and a reference(link) to the next node in the list.

The elements in a linked list are linked using pointers as shown in the below image:

![Linked List](https://media.geeksforgeeks.org/wp-content/cdn-uploads/gq/2013/03/Linkedlist.png)

## Types of Linked List

There are two types of linked list:

1. Singly Linked List
2. Doubly Linked List

There is also a variation of linked list called circular linked list.

## Singly Linked List

In singly linked list, each node contains a data field and a reference(link) to the next node in the list.

![Singly Linked List](https://media.geeksforgeeks.org/wp-content/cdn-uploads/gq/2013/03/Linkedlist.png)

## Doubly Linked List

In doubly linked list, each node contains a data field, and two reference(link) fields, one to the next node in the list and one to the previous node in the list.

![Doubly Linked List](https://media.geeksforgeeks.org/wp-content/cdn-uploads/gq/2014/03/DLL1.png)

## Circular Linked List

A circular linked list is a linked list where all nodes are connected to form a circle. There is no NULL at the end. A circular linked list can be a singly circular linked list or doubly circular linked list.

![Circular Linked List](https://media.geeksforgeeks.org/wp-content/cdn-uploads/gq/2014/03/circularlinkedlist.png)

## Advantages of Linked List

1. Dynamic size
2. Ease of insertion/deletion

## Disadvantages of Linked List

1. Random access is not allowed. We have to access elements sequentially starting from the first node. So we cannot do binary search with linked lists.
2. Extra memory space for a pointer is required with each element of the list.
3. Not cache friendly. Since array elements are contiguous locations, there is locality of reference which is not there in case of linked lists.

## Applications of Linked List

1. Implementation of stacks and queues
2. Implementation of graphs
3. Used in undo functionality at many places like editors, photoshop.
4. Used in Hash Table implementation
5. Used in Operating System for process management
6. Used in File System for file management
7. Used in Memory management
8. Used in Garbage Collection
9. Used in Implementing Adjacency List
10. Used in Implementing Adjacency Matrix
11. Used in Implementing Graphs
12. Used in Implementing Linked Representation of Graphs
13. Used in Implementing Sparse Matrix
14. Used in Implementing Polynomial Addition
15. Used in Implementing Polynomial Multiplication

## Summary

1. Linked list is a linear data structure, in which the elements are not stored at contiguous memory locations.
2. Each node contains a data field and a reference(link) to the next node in the list.
3. There are two types of linked list: singly linked list and doubly linked list.
4. There is also a variation of linked list called circular linked list.
5. Advantages of linked list: Dynamic size, Ease of insertion/deletion
6. Disadvantages of linked list: Random access is not allowed, Extra memory space for a pointer is required with each element of the list, Not cache friendly.
7. Applications of linked list: Implementation of stacks and queues, Implementation of graphs, Used in undo functionality at many places like editors, photoshop.

## Contributing

If you have any suggestions or improvements, feel free to open an issue or pull request.
