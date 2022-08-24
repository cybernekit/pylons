// Mocks generated by Mockito 5.0.14 from annotations
// in pylons_wallet/test/unit_testing/stores/wallet_store_imp_test.dart.
// Do not manually edit this file.

import 'dart:async' as i4;

import 'package:cosmos_utils/credentials_storage_failure.dart' as i5;
import 'package:dartz/dartz.dart' as i2;
import 'package:mockito/mockito.dart' as i1;
import 'package:transaction_signing_gateway/gateway/transaction_signing_gateway.dart';
import 'package:transaction_signing_gateway/model/account_lookup_key.dart'
    as i7;
import 'package:transaction_signing_gateway/model/transaction_signing_failure.dart'
    as i6;
import 'package:transaction_signing_gateway/transaction_signing_gateway.dart'
    as i3;

// ignore_for_file: avoid_redundant_argument_values
// ignore_for_file: avoid_setters_without_getters
// ignore_for_file: comment_references
// ignore_for_file: implementation_imports
// ignore_for_file: invalid_use_of_visible_for_testing_member
// ignore_for_file: prefer_const_constructors
// ignore_for_file: unnecessary_parenthesis
// ignore_for_file: camel_case_types

class _FakeEither_0<L, R> extends i1.Fake implements i2.Either<L, R> {}

/// A class which mocks [TransactionSigningGateway].
///
/// See the documentation for Mockito's code generation for more information.
class MockTransactionSigningGateway extends i1.Mock
    implements i3.TransactionSigningGateway {
  MockTransactionSigningGateway() {
    i1.throwOnMissingStub(this);
  }

  @override
  i4.Future<
      i2.Either<i5.CredentialsStorageFailure, i2.Unit>> updateAccountPublicInfo(
          {i3.AccountPublicInfo? info}) =>
      (super.noSuchMethod(
          Invocation.method(#updateWalletPublicInfo, [], {#info: info}),
          returnValue:
              Future<i2.Either<i5.CredentialsStorageFailure, i2.Unit>>.value(
                  _FakeEither_0<i5.CredentialsStorageFailure, i2.Unit>())) as i4
          .Future<i2.Either<i5.CredentialsStorageFailure, i2.Unit>>);

  @override
  i4.Future<i2.Either<i5.CredentialsStorageFailure, List<i3.AccountPublicInfo>>>
      getAccountsList() => (super.noSuchMethod(
          Invocation.method(#getWalletsList, []),
          returnValue:
              Future<i2.Either<i5.CredentialsStorageFailure, List<i3.AccountPublicInfo>>>.value(
                  _FakeEither_0<i5.CredentialsStorageFailure,
                      List<i3.AccountPublicInfo>>())) as i4
          .Future<i2.Either<i5.CredentialsStorageFailure, List<i3.AccountPublicInfo>>>);
  @override
  i4.Future<i2.Either<i6.TransactionSigningFailure, bool>> verifyLookupKey(
          i7.AccountLookupKey? walletLookupKey) =>
      (super.noSuchMethod(
              Invocation.method(#verifyLookupKey, [walletLookupKey]),
              returnValue:
                  Future<i2.Either<i6.TransactionSigningFailure, bool>>.value(
                      _FakeEither_0<i6.TransactionSigningFailure, bool>()))
          as i4.Future<i2.Either<i6.TransactionSigningFailure, bool>>);
}
