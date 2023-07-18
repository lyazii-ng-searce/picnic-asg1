const { ApolloServer, gql } = require('apollo-server');

// Sample data store
const names = [];

// GraphQL schema
const typeDefs = gql`
  type Name {
    id: ID!
    firstName: String!
    lastName: String!
  }

  type Query {
    names: [Name!]!
  }

  type Mutation {
    addName(firstName: String!, lastName: String!): Name!
  }
`;

// GraphQL resolvers
const resolvers = {
  Query: {
    names: () => names,
  },
  Mutation: {
    addName: (parent, args) => {
      const { firstName, lastName } = args;
      const id = String(names.length + 1);
      const newName = { id, firstName, lastName };
      names.push(newName);
      return newName;
    },
  },
};

// Create Apollo Server instance
const server = new ApolloServer({ typeDefs, resolvers });

// Start the server
server.listen().then(({ url }) => {
  console.log(`Server running at ${url}`);
});
