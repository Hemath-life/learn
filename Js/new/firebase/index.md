# Firebase Authentication with fireStore

To implement Firebase Authentication popups in a React application, follow these steps:

1. **Set up Firebase**: If you haven't already, create a Firebase project on the [Firebase Console](https://console.firebase.google.com/) and obtain your Firebase configuration object.

2. configure the authentication and firestore in firebase console .

3. **Install Firebase**: Install the Firebase SDK and React Firebase hooks by running the following command in your project directory:

   ```bash
   npm install firebase
   ```

## initialize Firebase with firestore:

Initialize Firebase in your application by creating a Firebase configuration file (e.g., firebase.js) and importing it where needed.

```JS
c// Import necessary Firebase functions and configurations
import { initializeApp } from "firebase/app";
import { getAuth, GoogleAuthProvider, signInWithRedirect } from "firebase/auth";

// Replace with your Firebase project configuration
const firebaseConfig = {
  apiKey: "YOUR_API_KEY",
  authDomain: "YOUR_AUTH_DOMAIN",
  projectId: "YOUR_PROJECT_ID",
  storageBucket: "YOUR_STORAGE_BUCKET",
  messagingSenderId: "YOUR_MESSAGING_SENDER_ID",
  appId: "YOUR_APP_ID",
};

// Initialize Firebase with your configuration
const firebaseApp = initializeApp(firebaseConfig);

// Create a GoogleAuthProvider instance
const googleProvider = new GoogleAuthProvider();

// Customize the sign-in prompt to 'select_account' (optional)
googleProvider.setCustomParameters({
  prompt: "select_account",
});

// Get the authentication instance
export const auth = getAuth();

// Function to sign in with Google using a redirect
export const signInWithGoogleRedirect = () => signInWithRedirect(auth, googleProvider);



// =======================
//  FIRST STORE SECTION
// =======================
// same file firebase.js

// initializes Firestore database.
export const db = getFirestore();


// Function to create or update a user profile document in Firestore
export const createUserProfileDocument = async (userAuth, additionalData) => {
  // Check if a user is authenticated
  if (!userAuth) return;

  // Define a reference to the user's document in Firestore
  const userDocRef = doc(db, 'users', userAuth.uid);

  // Retrieve the user's document snapshot from Firestore
  const userSnapshot = await getDoc(userDocRef);

  // Check if the user's document does not exist
  if (!userSnapshot.exists()) {
    // If the user document doesn't exist, create it with additionalData
    const { displayName, email } = userAuth;
    const createdAt = new Date();

    try {
      await setDoc(userDocRef, {
        // Customize this structure according to your user profile data
        displayName,
        email,
        createdAt,
        // ...other fields from additionalData
      });
    } catch (error) {
      console.error('Error creating user document:', error);
    }
  } else {
    // If the user document exists, update it with additionalData
    try {
      await setDoc(userDocRef, {
        // Merge existing data with additionalData
        ...userSnapshot.data(),
        updatedAt: new Date(), // Update the "updatedAt" field with the current date
        // ...other fields from additionalData that need updating
      }, { merge: true });
    } catch (error) {
      console.error('Error updating user document:', error);
    }
  }
};

```

## popup authentication (signInWithGooglePopup)

Create a React component that handles the Firebase Authentication popups. Here's an example component:

```JS
// AuthComponent.js
import {
  signInWithGooglePopup,
  createUserProfileDocument,
} from '../../utils/firebase/firebase.utils';

const SignIn = () => {
  const logGoogleUser = async () => {
    const { user } = await signInWithGooglePopup();
    createUserProfileDocument(user);
  };

  return (
    <div>
      <h1>Sign In Page</h1>
      <button onClick={logGoogleUser}>Sign in with Google Popup</button>
    </div>
  );
};

export default SignIn;

```

## redirect-based authentication (signInWithGoogleRedirect)

### issue in redirect

- this will redirect to the firstbase sing in application with different domain so it's unmounted react
- return comes with same local host domain so we lost the state and the react app mounted again the new

```jsx
// AuthComponent.js
import {
  signInWithGoogleRedirect,
  createUserProfileDocument,
} from "../../utils/firebase/firebase.utils";

const SignIn = () => {
  const logGoogleUser = async () => {
    // this will redirect to the firstbase sing in application with different domain so it's unmounted react
    // return comes with same local host domain so we lost the state and the react app mounted again the new
    const { user } = await signInWithGoogleRedirect();
    createUserProfileDocument(user);
  };

  return (
    <div>
      <h1>Sign In Page</h1>
      <button onClick={logGoogleUser}>Sign in with Google Popup</button>
    </div>
  );
};

export default SignIn;
```

## resolve this issue

## Firebase Authentication Flow

1. **Trigger Firebase Authentication**

   - When the "Sign In" button is clicked, it should trigger a Firebase authentication method (e.g., `signInWithRedirect` or `signInWithPopup`) that redirects the user to Firebase's authentication domain. This step will unmount your React app temporarily, but it's a necessary part of the authentication process.

2. **User Sign-In**

   - After the user completes the sign-in on the Firebase authentication page, they will be redirected back to your React app. Firebase will handle this redirect for you.

3. **Firebase Authentication State**

   - Firebase will store information about the authentication in its internal state. You typically don't need to worry about the specifics of this storage; Firebase handles it for you.

4. **Retrieve Authentication Result**

   - To retrieve the authentication result after the redirect, you should use the `getRedirectResult` method provided by Firebase Authentication. You pass `auth` as a parameter to this method. It will return an object containing user information if the user signed in successfully, or an error if there was a problem. You can handle the result in your React component.

5. **Use `useEffect` Hook**
   - You can use the `useEffect` hook to run code when your React component is mounted. In this case, you can use `useEffect` to handle the authentication result obtained in Step 4 and store user data in your React app's state or context if the authentication is successful.

```JSX
// AuthComponent.js

import { useEffect } from 'react';
import { getRedirectResult } from 'firebase/auth';
import {
  auth,
  signInWithGoogleRedirect,
  createUserProfileDocument,
} from "../../utils/firebase/firebase.utils";

const SignIn = () => {


  useEffect(()={
   const { user } = getRedirectResult(auth)
     createUserProfileDocument(user);
  },[])


  return (
    <div>
      <h1>Sign In Page</h1>
      <button onClick={signInWithGoogleRedirect}>Sign in with Google Popup</button>
    </div>
  );
};

export default SignIn;
```
