import Footer from '../components/umum/footer'
import Navbar from '../components/umum/navbar'
import styles from '../styles/Home.module.css'
import CariRiwayatTesComponent from '../components/cari-riwayat-tes/cari-riwayat-tes-component'

const CariRiwayatTes = () => {
    return(
        <main className={styles.main}>
            <Navbar />
            <CariRiwayatTesComponent />
            <Footer />
        </main>
    )
}

export default CariRiwayatTes;