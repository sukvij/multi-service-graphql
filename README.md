# multi-service-graphql

# Queries
query GetUser {
  getUser(id: "1") {
    id
    name
    email
    posts {
      id
      title
      content
    }
  }
}

query GetUsers {
  getUsers {
    id
    name
    email
    posts {
      id
      title
    }
  }
}

query GetPost {
  getPost(id: "1") {
    id
    title
    content
    user {
      id
      name
      email
    }
  }
}

query GetPosts {
  getPosts {
    id
    title
    content
    user {
      id
      name
    }
  }
}

# Mutations
mutation CreateUser {
  createUser(name: "John Doe", email: "john@example.com") {
    id
    name
    email
  }
}

mutation UpdateUser {
  updateUser(id: "1", name: "Jane Doe", email: "jane@example.com") {
    id
    name
    email
  }
}

mutation DeleteUser {
  deleteUser(id: "1")
}

mutation CreatePost {
  createPost(title: "My Post", content: "Hello, world!", userId: "1") {
    id
    title
    content
    user {
      id
      name
    }
  }
}

mutation UpdatePost {
  updatePost(id: "1", title: "Updated Post", content: "Updated content") {
    id
    title
    content
  }
}

mutation DeletePost {
  deletePost(id: "1")
}