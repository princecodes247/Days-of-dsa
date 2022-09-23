# Queue and Priority Queue

## Queue

A queue is a data structure that stores items in a First-In-First-Out (FIFO) manner. In a FIFO data structure, the first element added to the queue will be the first one to be removed. This is equivalent to the requirement that once a new element is added, all elements that were added before have to be removed before the new element can be removed.

A real-world example of a queue is the line of students waiting outside a classroom. The first student in line is the first one to enter the classroom, and the last student in line is the last one to enter. The first student to enter the classroom is also the first one to leave it, and so on.

## Priority Queue

A priority queue is a data structure that stores items in a First-In-First-Out (FIFO) manner, but with a twist. Each element of the queue has a priority associated with it. An element with high priority is dequeued before an element with low priority. If two elements have the same priority, they are served according to their order in the queue.

A real-world example of a priority queue is the line of students waiting outside a classroom. The first student in line is the first one to enter the classroom, and the last student in line is the last one to enter. The first student to enter the classroom is also the first one to leave it, and so on. However, if a teacher enters the classroom, all the students have to leave so that the teacher can teach. The teacher is served before any student, regardless of how long the student has been waiting.

## Uses of Queue

1.  Used in CPU scheduling, Disk Scheduling.
2.  Used in Graph algorithms like Breadth First Search, for file IO like when data is read from a file and needs to be processed line by line.
3.  Used in many problems on shortest path like Ford-Fulkerson algorithm.
4.  In real life scenario, Call Center phone systems uses Queues to hold people calling them in an order, until a service representative is free.
5.  Handling of interrupts in real-time systems.
6.  Used in async messaging where data is sent to be processed asynchronously.
7.  Used in Operating Systems. For example, when a process requests the OS for some memory, the OS searches for a free memory block in the main memory and places the process in a queue, until the memory becomes available.
8.  Used in Network packets. Packets are kept in a queue when they are waiting to be sent over a network.
9.  Used in BFS of Graph.
10. Used in Cache memory.

## Uses of Priority Queue

1.  Used in Graph algorithms like Dijkstra’s shortest path and Prim’s Minimum Spanning Tree.
2.  Used in scheduling algorithms like Job Scheduling and CPU Scheduling.
3.  Used in Huffman Coding.
4.  Used in Graph Coloring.
5.  Priority queues are implemented using heaps.
6.  Priority queues are used in Dijkstra’s algorithm for finding the shortest path in a graph.
7.  Priority queues are used in Huffman coding for compressing data.
8.  Priority queues are used in the Prim’s algorithm for finding the minimum spanning tree in a graph.
9.  Priority queues are used in the A\* search algorithm for finding the shortest path in a graph.
10. Priority queues are used in the Ford-Fulkerson algorithm for finding the maximum flow in a graph.
11. Priority queues are used in the Bellman-Ford algorithm for finding the shortest path in a graph.
12. Priority queues are used in the Kruskal’s algorithm for finding the minimum spanning tree in a graph.
13. Priority queues are used in the Floyd-Warshall algorithm for finding the shortest path in a graph.
14. Priority queues are used in the Johnson’s algorithm for finding the shortest path in a graph.
15. Priority queues are used in the D\* algorithm for finding the shortest path in a graph.

16. Used in Network packets. Packets are kept in a queue when they are waiting to be sent over a network.
17. Used in event management systems.
18. Used in Operating Systems. For example, when a process requests the OS for some memory, the OS searches for a free memory block in the main memory and places the process in a queue, until the memory becomes available.
19. Used in Cache memory.

## Summary

1.  Queue is a linear data structure which follows a particular order in which the operations are performed. The order is First In First Out (FIFO).
2.  A good example of queue is any queue of consumers for a resource where the consumer that came first is served first. The difference between stacks and queues is in removing. In a stack we remove the item the most recently added; in a queue, we remove the item the least recently added.
3.  The difference between a queue and a priority queue is that the priority queue removes the object with the highest priority first.

## Contributing

If you have any suggestions or improvements, feel free to open an issue or pull request.
