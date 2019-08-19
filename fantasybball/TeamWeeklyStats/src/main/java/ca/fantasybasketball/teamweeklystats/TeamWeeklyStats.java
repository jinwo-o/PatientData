package ca.fantasybasketball.teamweeklystats;
//
import ca.fantasybasketball.teamweeklystats.model.Player;
import ca.fantasybasketball.teamweeklystats.stats.PlayerWeeklyTotal;
import ca.fantasybasketball.teamweeklystats.readers.ReadCSV;
import ca.fantasybasketball.teamweeklystats.model.Roster;
import ca.fantasybasketball.teamweeklystats.stats.AddStats;
import ca.fantasybasketball.teamweeklystats.model.Stats;
import java.io.IOException;


//
public class TeamWeeklyStats {

    public static void main(String[] args) throws IOException {
        Roster roster = ReadCSV.CSV_Reader("/users/jin/work/fantasybball/FantasyCSV.csv");
        System.out.println("Yessir");
        Stats totalStats = new Stats();
        for(Player player : roster.getPlayers()) {
            Stats stats = PlayerWeeklyTotal.calculateWeeklyTotal(player);
            totalStats = AddStats.total(totalStats, stats);
        }

        System.out.println("Minutes: "+ totalStats.getMinutes());
        System.out.println("Threes: " + totalStats.getThrees());
        System.out.println("Rebounds: " + totalStats.getRebounds());
        System.out.println("Assists: " + totalStats.getAssists());
        System.out.println("Steals: " + totalStats.getSteals());
        System.out.println("Blocks: " + totalStats.getBlocks());
        System.out.println("Turnovers: " + totalStats.getTurnovers());
        System.out.println("DD: " + totalStats.getDD());
        System.out.println("Points: " + totalStats.getPoints());
    }

}
