import React from "react";
import { useState } from 'react'

const TambahPenyakitComponent = () => {
    const [nama, setNama] = useState("");
    const [file, setFile] = useState("");

    const tambahPenyakit = () => {
            // Check if document is finally loaded
            console.log(nama)
            if(!nama) {
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
    }

    return(
<       div
        id="intro-example"
        class="p-5 bg-image"
        >
        <div class="mask" style={{backgroundColor: "rgba(0, 0, 0, 0.2)",borderRadius: '4vh', paddingTop:"10vh", paddingBottom:"12vh"}}>
            <div class="d-flex justify-content-center align-items-center h-100">
            <div class="text-white">
                <h1 class="mb-3">Tambahkan Penyakit</h1>
                    <form class="align-items-left justify-content-left" onSubmit={tambahPenyakit}>
                        <div style={{marginTop:"5vh"}}>
                        <label for="formFile" class="form-label">Tambahkan Nama Penyakit</label>
                        <input onChange={(e) => setNama(e.target.value)} class="form-control" id="namaPenyakit" type="text" placeholder="Tambahkan Nama Penyakit" aria-label="default input example"></input>
                        </div>
                        <div class="mb-3" style={{marginTop:"3vh"}}>
                            <label for="formFile" class="form-label">Masukkan DNA</label>
                            <input size="100" class="form-control" type="file" id="formFile" onChange={(e) => setFile(e.target)}/>
                        </div>
                        <button type="submit" class="btn btn-dark btn-rounded" style={{borderRadius: "2vh", marginTop: "2vh", backgroundColor: "transparent", color: "white", borderColor: "white"}}>Tambahkan</button>
                    </form>
            </div>
            </div>
        </div>
        </div>
    )
}

module.exports = TambahPenyakitComponent;