import Image from 'next/image';
import styles from './page.module.css';

import turtleImg from 'public/turtle01.jpg';

import { IoChevronDown } from 'react-icons/io5';

export default function Home() {
  return (
    <>
      <div className={styles.hero_container}>
        <div className={styles.heroText_wrapper}>
          <h1>水平思考の間へようこそ！</h1>
          <p>「はい」か「いいえ」の質問をして、答を導け！</p>
        </div>
        <div className={styles.heroImg_wrapper}>
          <Image src={turtleImg} alt='turtle' layout='responsive' priority />
        </div>
      </div>
      <div className={styles.downText_wrapper}>
        <IoChevronDown />
        <h2>Scroll down...</h2>
      </div>
      <div className={styles.top_container}>
        <h1>問題一覧</h1>
        <h1>新着情報</h1>
      </div>
    </>
  );
}
