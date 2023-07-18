import React from 'react';
import { ApolloClient, ApolloProvider, InMemoryCache } from '@apollo/client';
import AddUser from './components/getUsers';

// Create an Apollo Client instance
const client = new ApolloClient({
  uri: 'http://localhost:4000/graphql', // Replace with your GraphQL server URL
  cache: new InMemoryCache(),
});

const App = () => {
  return (
    <ApolloProvider client={client}>
      <div>
        <h1>React App</h1>
        <AddUser />
      </div>
    </ApolloProvider>
  );
};

export default App;
