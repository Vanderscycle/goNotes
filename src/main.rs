#[allow(non_snake_case)]
mod cliParser;
fn main() {
    let receivedMsg: String = cliParser::parser("Hello, world!");
    println!("{}", &receivedMsg);
    println!(
        "Event simpler method {}",
        cliParser::parser("STRAIGHT FROM FN")
    );
}
