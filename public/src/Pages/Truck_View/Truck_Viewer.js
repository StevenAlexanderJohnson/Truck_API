import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import styles from "./Truck_Viewer.module.css";

function TruckViewer() {
    const navigator = useNavigate();
    const [truck_data, setTruck_data] = useState([]);
    const [qrCode, setQrCode] = useState(null);

    // Generate a list containing all the truck data
    const getTruckData = truck_data.map(truck => {
        return (
            <ul key={truck.Truck_Id}>
                <li>
                    <label htmlFor="truckId">ID: </label>
                    <span name="truckId">{truck.Truck_Id}</span>
                </li>
                <li>
                    <label htmlFor="truckPrice">Price: </label>
                    <span name="truckPrice">{truck.Truck_Price}</span>
                </li>
                <li>
                    <label htmlFor="truckNumber">Number: </label>
                    <span name="truckNumber">{truck.Truck_Number}</span>
                </li>
                <li>
                    <label htmlFor="truckType">Type: </label>
                    <span name="truckType">{truck.Truck_Type}</span>
                </li>
            </ul>
        )
    });

    // Generate an image to hold the QR code
    const qrCodeImage = qrCode ? <img src={qrCode} onClick={() => DownloadQrCode()} alt="QR Code" /> : null;

    function DownloadQrCode() {
        // Download the QR code as a PNG image
        const link = document.createElement("a");
        link.href = qrCode;
        link.download = "QR Code";
        link.click();
    }

    // Run on initial render
    useEffect(() => {
        // Get the split window location pathname
        const pathname = window.location.pathname.split('/');
        if (pathname.length > 4) {
            alert("Invalid URL");
            navigator("/");
        }
        // Get the truck data from the api
        fetch(`http://localhost:8080/api/${pathname[2]}`, {
            method: 'GET',
            credentials: 'include'
        })
            .then(res => {
                if (res.ok) {
                    return res.json()
                }
                else if (res.status === 401) {
                    console.log(res);
                    alert("Your session has expired. Please log in again.");
                    navigator("/login");
                    return;
                }

            })
            .then(data => {
                setTruck_data(data);
            });

        // Get the QR code from the api
        fetch(`http://localhost:8080/qr/${pathname[2]}`, {
            method: 'GET',
            credentials: 'include'
        })
            .then(res => res.blob())
            .then(blob => {
                const objectURL = URL.createObjectURL(blob);
                setQrCode(objectURL);
            })
            .catch(err => {
                console.log(err);
            });
    }, [navigator]);

    // JSX for the page
    return (
        <div className="FadeIn">
            <header className={styles.TruckContainer}>
                <input type="button" className={styles.BackButton + " YellowButton"} value="Back" onClick={() => navigator("/")} />
                {qrCodeImage}
                <h1 className="InvertedHeader">Truck Data</h1>
                {getTruckData}
            </header>
        </div>
    )
}

export default TruckViewer;