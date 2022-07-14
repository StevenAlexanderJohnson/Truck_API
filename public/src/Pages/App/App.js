import styles from "./App.module.css";
import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';

function App() {
    const navitagor = useNavigate();
    const [truckId, setTruckId] = useState('');

    useEffect(() => {
        // Ping the api to see if the user is logged in
        fetch('http://localhost:8080/', {
            method: 'GET',
            credentials: 'include'
        }).then(res => {
            // If the use is not logged in, redirect to the login page
            if (res.status === 401) {
                navitagor('/login');
            }
        });
    }, [navitagor])

    const logout = () => {
        fetch("http://localhost:8080/logout", {
            method: "POST",
            credentials: "include"
        }).then(res => {
            if (res.status === 200) {
                navitagor('/login');
            }
        });
    }

    const searchId = () => {
        if (truckId === '') {
            alert('Please enter a truck ID');
            return;
        }
        navitagor("/truck_view/" + truckId);
    }

    return (
        <div>
            <header className={styles.AppContainer + " FadeIn"}>
                <div>
                    <input type="button" className={styles.LogOutButton + " YellowButton"} value="Log Out" onClick={() => logout()} />
                </div>
                <h1 className="InvertedHeader">Welcome to the QR Demo</h1>
                <div className='VerticalContainer'>
                    <div className='SlideInputField'>
                        <input type="text" placeholder='Truck ID' id="truckId" className='SlideInput' onInput={e => setTruckId(e.target.value)} />
                        <label htmlFor="truckId" className='SlideLabel'>Search Truck ID</label>
                    </div>
                    <input type="button" className="YellowButton" value="Search" onClick={() => searchId()} />
                </div>
            </header>
        </div>
    );
}

export default App;
