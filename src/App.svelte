<script>
    import { FirebaseApp, User } from 'sveltefire'
    import firebase from 'firebase/app'
    import 'firebase/firestore'
    import 'firebase/auth'
    //import Chats from './components/Chats.svelte'
    import AuthForm from './components/AuthForm.svelte'
    let firebaseConfig = {
		apiKey: "AIzaSyDmlKBZoDQ-dWyeo9hCgZkv8ZSNPBKxY7Q",
		authDomain: "setec-cd5c7.firebaseapp.com",
		projectId: "setec-cd5c7",
		storageBucket: "setec-cd5c7.appspot.com",
		messagingSenderId: "64309028551",
		appId: "1:64309028551:web:117215c31b268f02da80e2",
		measurementId: "G-8RCN4JS31F"
    }
    firebase.initializeApp(firebaseConfig)
</script>

<style>
    .app {
        min-height: 100vh;
        display: flex;
        flex-direction: column;
        max-height: 100vh;
    }
</style>

<main class="container section app">

    <FirebaseApp {firebase}>
        <User let:user let:auth>
            <Chats {user} />
            <button class="button is-fullwidth" on:click={() => auth.signOut()}>Leave Chat</button>
            <div slot="signed-out">
                <AuthForm {auth} />
            </div>
        </User>

    </FirebaseApp>

</main>