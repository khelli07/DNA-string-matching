

const Navbar = () => {
    return(
        <div>
            <nav class="navbar navbar-expand-lg navbar-light">
                <a class="navbar-brand text-white" href="/">Informagila</a>
                <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarTogglerDemo02" aria-controls="navbarTogglerDemo02" aria-expanded="false" aria-label="Toggle navigation">
                    <span class="navbar-toggler-icon"></span>
                </button>

                <div class="collapse navbar-collapse" id="navbarTogglerDemo02">
                    <ul class="navbar-nav mt-2 mt-lg-0" style={{marginLeft: 'auto'}}>
                    <li class="nav-item active">
                        <a class="nav-link text-white" href="tambah-penyakit">Tambah Penyakit</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link text-white" href="tes-dna">Tes DNA</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link text-white" href="cari-riwayat-tes">Cari Riwayat Tes</a>
                    </li>
                    </ul>
                </div>
            </nav>
        </div>
    )
}

module.exports = Navbar;