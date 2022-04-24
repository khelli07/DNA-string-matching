import React from "react";
import { useState } from 'react'

const CariRiwayatTesComponent = () => {
    const [inputPengguna, setInputPengguna] = useState("");
    const [submitPenyakit, setSubmitPenyakit] = useState(false);
    const [dataPenyakitPengguna, setDataPenyakitPengguna] = useState([]);

    const cariDataPenyakit = () => {
            // Check if document is finally loaded
            if(!inputPengguna) {
                alert("Anda belum memasukkan apapun!");
                return false;
            }
            console.log(inputPengguna)
            setSubmitPenyakit(true);
            setDataPenyakitPengguna(["10 Maret 2021 - Dia - Halo - Tes", "11 Maret 2021 - Dia - Halo - Tes", "12 Maret 2021 - Dia - Halo - Tes"])
            return false;
    }

    return(
<       div
        id="intro-example"
        class="p-5 bg-image"
        >
        <div class="mask" style={{backgroundColor: "rgba(0, 0, 0, 0.2)",borderRadius: '4vh', paddingTop:"10vh", paddingBottom: submitPenyakit? '10vh' : '26vh'}}>
            <div class="d-flex justify-content-center align-items-center h-100">
            <div class="text-white">
                <h1 class="mb-3">Tes DNA</h1>
                <form class="align-items-left justify-content-left">
                    <div style={{marginTop:"5vh"}}>
                    <label for="formFile" class="form-label">Masukkan Tanggal, Penyakit, atau kombinasi keduanya</label>
                    <input onChange={(e) => setInputPengguna(e.target.value)} class="form-control" id="namaPengguna" type="text" placeholder="Masukkan Tanggal, Penyakit, atau kombinasi keduanya" aria-label="default input example"></input>
                    </div>
                    <button type="button" onClick={cariDataPenyakit} class="btn btn-dark btn-rounded" style={{borderRadius: "2vh", marginTop: "4vh", backgroundColor: "transparent", color: "white", borderColor: "white"}}>Identifikasi</button>
                </form>
                <h3 class="mb-3" hidden={!submitPenyakit} style={{marginTop: "4vh"}}>Hasil: </h3>
                <hr hidden={!submitPenyakit} style={{backgroundColor: 'white'}}></hr>
                {dataPenyakitPengguna.map((data) => {
                    return(
                        <div>
                            <p hidden={!submitPenyakit}>{data}</p>
                        </div>
                    )
                })}
            </div>
            </div>
        </div>
        </div>
    )
}

module.exports = CariRiwayatTesComponent;