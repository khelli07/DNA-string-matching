import Footer from '../components/umum/footer'
import Navbar from '../components/umum/navbar'
import styles from '../styles/Home.module.css'
import TesDNAComponent from '../components/tes-dna/tes-dna-component'

const TesDNA = () => {
    return(
        <main className={styles.main}>
            <Navbar />
            <TesDNAComponent />
            <Footer />
        </main>
    )
}

export default TesDNA;