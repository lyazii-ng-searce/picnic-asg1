import React, { useState } from 'react';
import { useMutation } from '@apollo/react-hooks';
import gql from 'graphql-tag';

const ADD_USER_MUTATION = gql`
  mutation AddUser($firstName: String!, $lastName: String!) {
    addUser(firstName: $firstName, lastName: $lastName) {
      
      id
      firstName
      lastName
    
    }
  }
`;

const AddUser = () => {
  const [firstName, setFirstName] = useState('');
  const [lastName, setLastName] = useState('');

  const [addUser, { loading, error }] = useMutation(ADD_USER_MUTATION);

  const handleSubmit = (e) => {
    e.preventDefault();

    addUser({ variables: { firstName, lastName } })
      .then(({ data }) => {
        // Handle success response
        console.log('Response:', data);
      })
      .catch((error) => {
        // Handle error response
        console.error('Error:', error);
      });

    // Clear form fields
    setFirstName('');
    
    setLastName('');
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="text"
        value={firstName}
        onChange={(e) => setFirstName(e.target.value)}
        placeholder="First Name"
      />
      <input
        type="text"
        value={lastName}
        onChange={(e) => setLastName(e.target.value)}
        placeholder="Last Name"
      />
      <button type="submit" disabled={loading}>
        Submit
      </button>
      {error && <p>Error: {error.message}</p>}
    </form>
  );
};

export default AddUser;
