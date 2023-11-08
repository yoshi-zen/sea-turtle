'use client';

import WrapperShadowBox from '@/components/wrapper-shadow-box';
import styles from './page.module.css';
import MyInput from '@/components/component-input';
import { useState } from 'react';

export default function Page() {
  const [newChecker, setNewChecker] = useState(false);

  const handleClick = () => {
    setNewChecker(!newChecker);
  };

  return (
    <>
      <WrapperShadowBox alignCenter style={{ width: '80%', maxWidth: '640px' }}>
        <h3>{newChecker ? '新規登録' : 'ログイン'}</h3>
        <MyInput type={'text'} label={'ユーザー名 / メールアドレス'} />
        <MyInput type={'password'} label={'パスワード'} />
        <p style={{ textAlign: 'center' }} onClick={handleClick}>
          {newChecker ? 'ログイン' : '新規登録'}はこちらから
        </p>
        <div className={styles.button_wrapper}>
          <p>パスワードをお忘れの場合</p>
          <button className={styles.button}>ログイン</button>
        </div>
      </WrapperShadowBox>
    </>
  );
}
