import React from "react";
import { useState } from 'react'

const TesDNAComponent = () => {
    const [namaPenyakit, setNamaPenyakit] = useState("");
    const [namaPengguna, setNamaPengguna] = useState("");
    const [submitPenyakit, setSubmitPenyakit] = useState(false);
    const [dataPenyakitPengguna, setDataPenyakitPengguna] = useState([]);
    const [file, setFile] = useState("");

    const cariDataPenyakitPengguna = () => {
            // Check if document is finally loaded
            console.log(namaPengguna)
            console.log(namaPenyakit)
            if(!namaPengguna) {
                alert("Anda belum memasukkan nama pengguna!");
                return;
            }
            if(!namaPenyakit) {
                alert("Anda belum memasukkan nama penyakit!");
                return;
            }
            if (!file) {
                alert("Anda belum memasukkan file apapun!");
            }
            else if (!file.files) {
                alert("File yang anda masukkan tidak disuppor! Hanya masukkan file berekstensi .txt");
            }
            else if (!file.files[0]) {
                alert("Anda belum memasukkan file apapun!");
            }
            else {
                var temp = file.files[0];

                var reader = new FileReader();
                var hasil = "";
                reader.onload = function(progressEvent){
                  var lines = this.result.split('\n');
                  hasil = lines[0];
                  console.log(hasil)
                    
                };
                reader.readAsText(temp);
            }
            setSubmitPenyakit(true);
            setDataPenyakitPengguna("10 Maret 2021 - Dia - Halo - Tes")
            return false;
    }

    return(
<       div
        id="intro-example"
        class="p-5 bg-image"
        >
        <div class="mask" style={{backgroundColor: "rgba(0, 0, 0, 0.2)",borderRadius: '4vh', paddingTop:"10vh", paddingBottom:"12vh"}}>
            <div class="d-flex justify-content-center align-items-center h-100">
            <div class="text-white">
                <h1 class="mb-3">Tes DNA</h1>
                <form class="align-items-left justify-content-left" onSubmit={cariDataPenyakitPengguna}>
                    <div style={{marginTop:"5vh"}}>
                    <label for="formFile" class="form-label">Masukkan Nama Pengguna</label>
                    <input onChange={(e) => setNamaPengguna(e.target.value)} class="form-control" id="namaPengguna" type="text" placeholder="Masukkan Nama Pengguna" aria-label="default input example"></input>
                    </div>
                    <div class="mb-3" style={{marginTop:"3vh"}}>
                        <label for="formFile" class="form-label">Masukkan DNA</label>
                        <input class="form-control" type="file" id="formFile" onChange={(e) => setFile(e.target)}/>
                    </div>
                    <div style={{marginTop:"3vh"}}>
                    <label for="formFile" class="form-label">Masukkan Nama Penyakit</label>
                    <input onChange={(e) => setNamaPenyakit(e.target.value)} class="form-control" id="namaPenyakit" type="text" placeholder="Masukkan Nama Penyakit" aria-label="default input example"></input>
                    </div>
                    <button type="button" onClick={cariDataPenyakitPengguna} class="btn btn-dark btn-rounded" style={{borderRadius: "2vh", marginTop: "2vh", backgroundColor: "transparent", color: "white", borderColor: "white"}}>Identifikasi</button>
                </form>
                <h3 class="mb-3" hidden={!submitPenyakit} style={{marginTop: "4vh"}}>Hasil: </h3>
                <hr hidden={!submitPenyakit} style={{backgroundColor: 'white'}}></hr>
                <p hidden={!submitPenyakit}>{dataPenyakitPengguna}</p>
            </div>
            </div>
        </div>
        </div>
    )
}

module.exports = TesDNAComponent;