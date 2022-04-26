import React from "react";
import axios from "axios";
import { BACKEND_URL } from "../../constant";
import { useState } from 'react'

const TambahPenyakitComponent = () => {
    const [nama, setNama] = useState("");
    const [file, setFile] = useState("");
    const [hasilRead, setHasilRead] = useState("");

    const tambahPenyakit = async (e) => {
            // Check if document is finally loaded
            e.preventDefault()
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
                reader.onload = async function(progressEvent){
                  var lines = this.result.split('\n');
                  hasil = lines[0];
                  setHasilRead(hasil);
                  console.log(hasil);
                  try {
                    const newDataPenyakit = {
                      name: nama,
                      pattern: hasil,
                    }
                    console.log(newDataPenyakit);
                    alert("Penyakit berhasil ditambahkan!");
                    axios.post(`http://127.0.0.1:8080/penyakit/new`, {
                        name: nama,
                        pattern: hasil,
                      })
                      .then((response) => {
                        console.log(response);
                        alert("Penyakit berhasil ditambahkan!");
                      }, (error) => {
                        console.log(error);
                      });
                      alert("Penyakit berhasil ditambahkan!");
/*                     const attempt = await axios({
                      method: "post",
                      url: `${BACKEND_URL}/penyakit/new`,
                      headers: {
                        "Content-Type": "application/json; charset=utf-8",
                      },
                      data: newDataPenyakit
                    }); */
                  }catch(err) {
                    alert(err.toString());
                  };
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