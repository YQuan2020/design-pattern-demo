package structural

import (
	"fmt"
	"log"
)

type Account struct {
	name string
}

func (a *Account) checkAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("account name is incorrect")
	}
	fmt.Println("account verified")
	return nil
}

func newAccount(accountName string) *Account {
	return &Account{
		name: accountName,
	}
}

type Wallet struct {
	balance int
}

func newWallet() *Wallet {
	return &Wallet{balance: 0}
}

func (w *Wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Println("wallet balance added successfully")
	return
}

func (w *Wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("balance is not sufficient")
	}
	fmt.Println("wallet balance is sufficient")
	w.balance = w.balance - amount
	return nil
}

type SecurityCode struct {
	code int
}

func (s *SecurityCode) checkCode(incomingCode int) error {
	if s.code != incomingCode {
		return fmt.Errorf("security code is incorrect")
	}
	fmt.Println("security code verified")
	return nil
}

func newSecurityCode(code int) *SecurityCode {
	return &SecurityCode{
		code: code,
	}
}

type Notification struct {
}

func (n *Notification) sendWalletCreditNotification() {
	fmt.Println("sending wallet credit notification")
}

func (n *Notification) sendWalletDebitNotification() {
	fmt.Println("sending wallet debit notification")
}

type Ledger struct {
}

func (l *Ledger) makeEntry(accountId, txnType string, amount int) {
	fmt.Printf("make ledger entry for accountId %s with txnType: %s for amount %d\n", accountId, txnType, amount)
	return
}

type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	securityCode *SecurityCode
	notification *Notification
	ledger       *Ledger
}

func newWalletFacade(accountId string, code int) *WalletFacade {
	fmt.Println("starting create account")
	walletFacade := &WalletFacade{
		account:      newAccount(accountId),
		securityCode: newSecurityCode(code),
		wallet:       newWallet(),
		notification: &Notification{},
		ledger:       &Ledger{},
	}
	fmt.Println("account created")
	return walletFacade
}

func (w *WalletFacade) addMoneyToWallet(accountId string, securityCode int, amount int) error {
	fmt.Println("starting add money to wallet")
	err := w.account.checkAccount(accountId)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountId, "credit", amount)
	return nil
}

func (w *WalletFacade) deductMoneyFromWallet(accountId string, securityCode int, amount int) error {
	fmt.Println("start debit money from wallet")
	err := w.account.checkAccount(accountId)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}
	w.notification.sendWalletDebitNotification()
	w.ledger.makeEntry(accountId, "debit", amount)
	return nil
}

func DoFacade() {
	walletFacade := newWalletFacade("abc", 1234)
	fmt.Println()

	err := walletFacade.addMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
