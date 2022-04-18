import Footer from '../components/umum/footer'
import Navbar from '../components/umum/navbar'
import styles from '../styles/Home.module.css'
import TambahPenyakitComponent from '../components/tambah-penyakit/tambah-penyakit-component'

const TambahPenyakit = () => {
    return(
        <main className={styles.main}>
            <Navbar />
            <TambahPenyakitComponent />
            <Footer />
        </main>
    )
}

export default TambahPenyakit;