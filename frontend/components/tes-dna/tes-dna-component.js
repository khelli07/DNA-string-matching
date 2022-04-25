import React from "react";
import { useState } from 'react';
import axios from "axios";
import { BACKEND_URL } from "../../constant";

const TesDNAComponent = () => {
    const [namaPenyakit, setNamaPenyakit] = useState("");
    const [namaPengguna, setNamaPengguna] = useState("");
    const [submitPenyakit, setSubmitPenyakit] = useState(false);
    const [dataPenyakitPengguna, setDataPenyakitPengguna] = useState([{date: "a", name: "a", disease: "a", result: "a", percentage: "a"}]);
    const [dataId, setDataId] = useState(-1);

    const [hasilRead, setHasilRead] = useState("");
    const [file, setFile] = useState("");

    const cariDataPenyakitPengguna = async (e) => {
            // Check if document is finally loaded
            e.preventDefault()
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
                reader.onload = async function(progressEvent){
                  var lines = this.result.split('\n');
                  hasil = lines[0];
                  try {
                    const newDataPenyakit = {
                        name: namaPengguna,   
                        sequence: hasil,
                        disease: namaPenyakit,
                    }
                    
                    const attempt = await axios({
                      method: "post",
                      url: `${BACKEND_URL}/diagnosis/new`,
                      headers: {
                        "Content-Type": "application/json",
                      },
                      data: newDataPenyakit
                    }).then(res => {
                        console.log(res.data.new_id);
                        setDataId(res.data.new_id);
                        fetchData(res.data.new_id - 1);
                        setSubmitPenyakit(true);
                    });

                  }catch(err) {
                    alert(err.toString());
                  };
                };
                reader.readAsText(temp);
            }
            return false;
    }

/*     React.useEffect(() => {
        const getDataUser = async () => {
          const { data: penyakitData } = await axios.get(`${BACKEND_URL}/${dataId}`, {
          } )
          setDataPenyakitPengguna(penyakitData);
        }
    }, []); */

    async function fetchData(id) {
        if(id > -1) {
            const { data: penyakitData } = await axios.get(
                `${BACKEND_URL}/diagnosis`,
            )
            console.log(penyakitData);
            let data = Object.values(penyakitData);
            setDataPenyakitPengguna({
                date: data[id].date,
                name: data[id].name,
                disease: data[id].disease,
                result: data[id].result,
                percentage: data[id].percentage,
            });
        }
      }

    React.useEffect(() => {
        fetchData(dataId);
    }, []);

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
                {dataId !== -1? <p hidden={!submitPenyakit}>{dataPenyakitPengguna.date} - {dataPenyakitPengguna.name} - {dataPenyakitPengguna.disease} - {dataPenyakitPengguna.result? "True" : "False"} - {dataPenyakitPengguna.percentage}</p> : <p></p>}
            </div>
            </div>
        </div>
        </div>
    )
}

module.exports = TesDNAComponent;