import styles from '../styles/Home.module.css'
import React from "react";

import Footer from '../components/umum/footer'
import Navbar from '../components/umum/navbar'
import Jumbotron from '../components/index/jumbotron';

export default function Home() {
  return (
    <div>
      <main className={styles.main}>
        <Navbar />
        <Jumbotron />
        <Footer />
      </main>
    </div>
  )
}
