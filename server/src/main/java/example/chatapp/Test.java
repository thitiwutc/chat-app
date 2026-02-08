package example.chatapp;

import java.time.LocalDateTime;
import java.util.List;
import lombok.extern.slf4j.Slf4j;
import org.springframework.boot.ApplicationArguments;
import org.springframework.boot.ApplicationRunner;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Component;

@Component
@Slf4j
public class Test implements ApplicationRunner {
  private final JdbcTemplate jdbcTemplate;

  public Test(JdbcTemplate jdbcTemplate) {
    this.jdbcTemplate = jdbcTemplate;
  }

  @Override
  public void run(ApplicationArguments args) throws Exception {
    List<LocalDateTime> result =
        this.jdbcTemplate.queryForList("SELECT NOW();", LocalDateTime.class);
    log.info("Query result: {}", result);
  }
}
