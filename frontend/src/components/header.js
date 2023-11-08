'use client';

import styles from '@/styles/header.module.css';
import Container from './container';
import { useRouter } from 'next/navigation';

export default function Header() {
  const router = useRouter();
  return (
    <header className={styles.header}>
      <Container large>
        <div className={styles.flexContainer}>
          <div className={styles.header_title} onClick={() => router.push('/')}>
            <h2>ウミガメのスープ</h2>
            <p>～水平思考の間～</p>
          </div>
          <ul className={styles.header_menu}>
            <li onClick={() => router.push('/')}>ホーム</li>
            <li onClick={() => router.push('/questions')}>問題一覧</li>
            <li onClick={() => router.push('/auth')}>log in</li>
          </ul>
        </div>
      </Container>
    </header>
  );
}
