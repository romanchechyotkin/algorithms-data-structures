use std::fmt::Display;

struct Queue<'a>  {
    head: Option<&'a Node<'a>>,
    tail: Option<&'a Node<'a>>,
}

struct Node<'a> {
    val: usize,
    next: Option<&'a Node<'a>>,
}



impl Display for Node<'_> {

    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let val = self.val;
        return match &self.next {
            Some(node) => write!(f, "{}; next {}", val, node.val),
            None => write!(f, "{}; NONE", self.val),
        } 
    }

}

impl<'a> Queue<'a> {

    fn new() -> Self {
        Queue {head: None, tail: None}
    }

    fn push(&mut self, value: usize) {
        let new_node = Node { val: value, next: None };

        match self.tail {
            None => {
                self.head = Some(&new_node);
                self.tail = Some(&new_node);
            }
            Some(rear_node) => {
                self.tail = Some(&new_node);
                rear_node.next = Some(&new_node);
            }
        }
    }

}

fn main() {
    let mut queue = Queue::new();

    queue.push(1);
    queue.push(2);
    queue.push(3);
    queue.push(4);
} 