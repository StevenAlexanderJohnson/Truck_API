import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import styles from './Login.module.css';
function Login() {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const logPassword = () => {
        console.log(password);
    }

    const navigator = useNavigate();

    const loginUser = () => {
        fetch("http://localhost:8080/sign_in", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                email: username,
                password: password
            }),
            credentials: "include"
        }).then(res => {
            if (res.status === 200) {
                setTimeout(() => {
                    navigator('/');
                }, 250)
                document.getElementById("container").classList.add("FadeOut");
            }
        });
    }

    return (
        <div id="container" className={styles.LoginContainer}>
            <header className={styles.LoginHeader}>
                <h1>Login</h1>
                <div className="VerticalContainer">
                    <div className="SlideInputField">
                        <input type="text" placeholder='Email' id="username" className='SlideInput' value={username} onInput={e => setUsername(e.target.value)} />
                        <label htmlFor="username" className='SlideLabel'>Email</label>
                    </div>
                    <div className="SlideInputField">
                        <input type="password" placeholder='Password' id="password" className='SlideInput' value={password} onInput={e => setPassword(e.target.value)} />
                        <label htmlFor="password" className='SlideLabel'>Password</label>
                    </div>
                </div>
                <div className={styles.ButtonContainer}>
                    <button className="YellowButton" onClick={() => loginUser()}>Login</button>
                    <button className="YellowButton" onClick={() => logPassword()}>Register</button>
                </div>
            </header>
        </div>
    );
}

export default Login;
