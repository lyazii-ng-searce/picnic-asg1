const { ApolloServer, gql } = require('apollo-server');

// Sample data store
const users = [];

// GraphQL schema
const typeDefs = gql`
  type User {
    id: ID!
    firstName: String!
    lastName: String!
  }

  type Query {
    users: [User!]!
  }

  type Mutation {
    addUser(firstName: String!, lastName: String!): User!
  }
`;

// GraphQL resolvers
const resolvers = {
  Query: {
    users: () => users,
  },
  Mutation: {
    addUser: (parent, args) => {
      const { firstName, lastName } = args;
      const id = String(users.length + 1);
      const newUser = { id, firstName, lastName };
      users.push(newUser);
      return newUser;
    },
  },
};

// Create Apollo Server instance
const server = new ApolloServer({ typeDefs, resolvers });

// Start the server
server.listen().then(({ url }) => {
  console.log(`Server running at ${url}`);
});
