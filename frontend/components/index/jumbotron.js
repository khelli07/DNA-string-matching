

const Jumbotron = () => {
    return(
        <div
        id="intro-example"
        class="p-5 text-center bg-image"
        >
        <div class="mask" style={{backgroundColor: "rgba(0, 0, 0, 0.2)", paddingTop: '20vh', paddingBottom: '20vh', borderRadius: '4vh'}}>
            <div class="d-flex justify-content-center align-items-center h-100">
            <div class="text-white">
                <h1 class="mb-3">DNA Pattern Matching</h1>
                <hr style={{backgroundColor: 'white'}}></hr>
                <h5 class="mb-4">
                    Website pencocokan Penyakit dengan DNA
                </h5>
                <a
                class="btn btn-outline-light btn-lg m-2"
                href="tambah-penyakit"
                role="button"
                >Tambahkan Penyakit</a>
                <a
                class="btn btn-outline-light btn-lg m-2"
                href="tes-dna"
                >Tes DNA</a>
            </div>
            </div>
        </div>
        </div>
    )
}

module.exports = Jumbotron;